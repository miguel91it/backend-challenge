# Golioth Challange

Project developed by João Miguel to solve the Golioth Challenge.

Following it will be described how the project was structured and hot to run it.

## Project Structure

The project is a monorepo containing a `src/ folder wich has:

    * main.go file               -> responsible to start the gateway application
    * Dockerfile file            -> responsible to build the gateway container image
    * gateway/ folder            -> responsible for group all files related to the gateway application
    * devices-simulation/ folder -> responsible for group all files related to the devices simulation application
    * docker-compose.yaml file   -> responsible for launch all the 3 containers (gateway, devices and mongodb)
    * .env file                  -> responsible for provide environment variables for the gateway and devices containers

Always you want to run the gateway application and the devices simulation, you'll use the `docker-compose` tool. To see how to run the solution, go to the end of this file.

## Gateway Module

Gateway Module is composed by 3 files:

    * gateway.go  -> responsible for launch the gateway server
    * handlers.go -> responsible for group the API endpoints handlers
    * weather.go  -> responsible for modeling the weather entity

The Weather Telemetry API has the following entity (model):

    * Weather

### Weather Entity

This entity groups data related to the environment weather sent by the devices.

    * Device_id           (string)  -> MAC Addres as device id
    * Timestamp           (int)     -> Timestamp of the measure
    * SoilMoisture        (float64) -> Rate of soil moisture (must be between 0 and 100%)
    * ExternalTemperature (float64) -> External temperature in celsius degree
    * ExternalHumidity    (float64) -> Rate of external humidity (must be between 0 and 100%)

## Devices Module

Devices Module is composed by 3 main files:

    * devices.go   -> contains a function to provides a list of MAC Address to the simualtion
    * telemetry.go -> contains the functions responsible for sending telemetry
    * main.go      -> responsible for coordinate the creation of one goroutine for each device and for the interval of sending management during the time of simulation desired.


## How to run the project

The project is containerized using a Golang image and for run it you must follow the next steps:

* First of all, clone the repository into your machine
> git clone https://github.com/miguel91it/backend-challenge.git

* Next you must enter in the root folder:
> cd backend-challenge

* Once inside at the root folder of the project, you must build and run all the containers (mongodb, gateway and devices) using docker compose:
> sudo docker-compose up --build gateway devices

The command above will build and run all the 3 containers but only gateway and devices containers will be attached and logged into standard output. This way you'll be able to see both logs during the gateway and devices containers executions.

* After the succesfull conclusion of the above command, you'll be able to attach to the gateway container and see its logs:
> sudo docker attach gateway

* Finally, to stop and remove all containers
> sudo docker-compose rm -sfv


## How to test the Gateway Server

It's possible to test manualy the telemetry endpoint of the API using command line tools like Curl or GUI tools like Postman.

Following we have some example of how to send a telemetry data to the gateway using `Curl`:

> POST: localhost:28000/api/v1/WeatherTelemetry

```
    Example of Request body:

        {
            "id":              "1",
            "timestamp":       12345678,
            "soil_moisture":   10,
            "ext_temperature": 10.1,
            "ext_humidity":    1.3
	    }
```

Full `Curl` command:

> curl --request POST 'localhost:28000/api/v1/WeatherTelemetry' \ \
--header 'Content-Type: application/json' \ \
--data-raw '{ \
		"id":              "02:42:ac:17:00:03", \
		"timestamp":       1629855252, \
		"soil_moisture":   10,\
		"ext_temperature": 10.1,\
		"ext_humidity":    1.3\
	}'

Experiment change soil_moisture, ext_temperature and ext_humidity values.