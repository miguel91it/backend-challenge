package gateway

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

var db Repository

type IGateway interface {
	Run()
}

type Gateway struct {
	PORT int
}

func NewGateway(port int, database Repository) (*Gateway, error) {

	if port < 9999 {
		return nil, fmt.Errorf("PORT must be greater than 10000")
	}

	db = database

	return &Gateway{
		PORT: port,
	}, nil
}

func (g *Gateway) Run() error {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/WeatherTelemetry", WeatherTelemetry).Methods("POST")

	addr := fmt.Sprintf("localhost:%v", g.PORT)

	fmt.Printf("Starting Gateway at %s\n", addr)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", g.PORT), router); err != nil {

		return fmt.Errorf("unable to raise the Gateway at the address %s. %v", addr, err)
	}

	return nil

}
