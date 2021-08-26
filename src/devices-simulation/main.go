package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Printf("Simulating Devices\n\n")

	// pointers to the parameters taken from command line flags
	numDevicesPtr := flag.Int("num-devices", 1, "an int")
	durationPtr := flag.Int("duration", 3, "an int")
	intervalPtr := flag.Int("interval", 1, "an int")

	// parse the command line flags
	flag.Parse()

	// getting the values of the flags pointers
	numDevices := *numDevicesPtr
	duration := time.Duration(*durationPtr) * time.Second
	interval := time.Duration(*intervalPtr) * time.Second

	fmt.Printf("This simulation will occur:\n\t-> for %d devices\n\t-> during %v at the total\n\t-> with %v of interval between sents\n\n", numDevices, duration, interval)

	// wait grouop for wait all the goroutines finish
	var wg sync.WaitGroup

	// get the first numDevices from the list of fake MAC Addresses
	devices, err := TopMAC(numDevices)

	if err != nil {

		fmt.Printf("%s", err)

		return

	}

	// create a slice of boolean cahnnels to send a terminate signal to the goroutines
	var dones []chan bool

	startTime := time.Now()

	// runs through devices list and raise a goroutine for each of them
	for _, device := range devices {

		// create a boolean channel to be passed into goroutine
		done := make(chan bool)

		// add a new goroutine to the wait group
		wg.Add(1)

		// raise a goroutine to send telemetry data
		go func(device string, interval time.Duration, done chan bool) {

			n := 0

			// create a ticker to control the interval between sents
			ticker := time.NewTicker(interval)

			// call the ticker stop at the end of the goroutine
			defer ticker.Stop()

			for {
				select {

				// arrived the time to send telemetry data (ticker posted a signal into it's channel 'C')
				case <-ticker.C:

					if err := sendTelemetry(createPayload(device, interval.Seconds(), n)); err != nil {

						fmt.Printf("%s\n", err)

					}

					n++

				// arrived the time to finish the goroutine
				case <-done:
					fmt.Printf("\nDevice %s finished", device)

					// tell to the wait group that this goroutine finished
					wg.Done()

					return
				}
			}

		}(device, interval, done)

		// add to the slice of dones channel each done channel created
		dones = append(dones, done)
	}

	// force synchronous wait for 'duration' time
	time.Sleep(duration)

	// send the signal to the goroutines finish itselves
	for _, done := range dones {
		done <- true

		// close the done channel
		close(done)

	}

	// force wait for all goroutines
	wg.Wait()

	// compute the entire duration of sent
	period := time.Since(startTime)

	fmt.Printf("\nDuration of the ingestion: %s\n", period)

}
