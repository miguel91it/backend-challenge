package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

type Weather struct {
	Device_id           string    `json:"id"`
	Timestamp           int       `json:"timestamp"`
	Datetime            time.Time `json:"datetime"`
	SoilMoisture        float64   `json:"soil_moisture"`
	ExternalTemperature float64   `json:"ext_temperature"`
	ExternalHumidity    float64   `json:"ext_humidity"`
}

// TODO:
func NewWeatherFromJson(jsonDecoder *json.Decoder) (*Weather, error) {

	var weather Weather

	if err := jsonDecoder.Decode(&weather); err != nil {

		return &Weather{}, fmt.Errorf("cannot decode json weather telemetry into Weather Model: %+v", err)
	}

	if err := weather.validateDeviceId(); err != nil {

		return &Weather{}, fmt.Errorf("device ID provided doesn't have valid ID: %+v", err)
	}

	if err := weather.validateSoilMoisture(); err != nil {

		return &Weather{}, fmt.Errorf("soil Moisture provided is not valid: %+v", err)
	}

	if err := weather.validateExtHumidity(); err != nil {

		return &Weather{}, fmt.Errorf("external Humidity provided is not valid: %+v", err)
	}

	return &weather, nil

}

func (w *Weather) validateDeviceId() error {
	if w.Device_id == "" {

		return fmt.Errorf("empty device ID provided: '%s'", w.Device_id)
	}

	return nil
}

func (w *Weather) validateSoilMoisture() error {

	if w.SoilMoisture > 100.0 || w.SoilMoisture < 0.0 {

		return fmt.Errorf("soil Moisture out of range [0 - 100]: %.2f", w.SoilMoisture)
	}

	return nil
}

func (w *Weather) validateExtHumidity() error {
	if w.ExternalHumidity > 100.0 || w.ExternalHumidity < 0.0 {

		return fmt.Errorf("external Humidity out of range [0 - 100]: %.2f", w.SoilMoisture)
	}

	return nil
}
