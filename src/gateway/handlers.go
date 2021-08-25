package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WeatherTelemetry(w http.ResponseWriter, r *http.Request) {

	// create weather model from json posted
	weather, err := NewWeatherFromJson(json.NewDecoder(r.Body))

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error receiving the weather data. %v", err)

		return
	}

	fmt.Printf("\nweather: %+v", weather)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Weather Data received succesfuly")
}
