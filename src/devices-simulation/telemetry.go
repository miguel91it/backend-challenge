package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func sendTelemetry(mapIn map[string]interface{}) error {

	host := os.Getenv("GATEWAY_HOST")
	port := os.Getenv("GATEWAY_PORT")
	path := "api/v1/WeatherTelemetry"
	uri := fmt.Sprintf("http://%v:%v/%v", host, port, path)

	body, err := json.Marshal(mapIn)

	if err != nil {

		return fmt.Errorf("error creating json to send to gateway\n %+v\n%+v", mapIn, err)
	}

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(body))

	if err != nil {

		return fmt.Errorf("error trying to send data to gateway\n %+v\n%+v", resp, err)
	}

	if resp.StatusCode != http.StatusOK {

		b, _ := io.ReadAll(resp.Body)

		return fmt.Errorf("resp Status: %+v\nResp Body: %s\nPayload: %+v", resp.Status, string(b), mapIn)
	}

	return nil
}

func createPayload(deviceId string) map[string]interface{} {

	return map[string]interface{}{
		"id":              deviceId,
		"timestamp":       time.Now().Unix(),
		"soil_moisture":   7.5,
		"ext_temperature": 10.1,
		"ext_humidity":    98.98,
	}

}
