package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWeatherModel(t *testing.T) {

	// weather_mock := map[string]interface{}

	mapD := map[string]interface{}{
		"id":              "1",
		"timestamp":       time.Now().Unix(),
		"soil_moisture":   7.5,
		"ext_temperature": 10.1,
		"ext_humidity":    11,
	}

	jsonWeather, err := mockJsonWeather(mapD)

	if err != nil {

		fmt.Printf("test case failed because there is something wrong creating Wetaher Mock Json to be tested: %v", err)

		return
	}

	// fmt.Println(string(jsonWeather))

	// fmt.Println(reflect.TypeOf(jsonWeather))

	// fmt.Println(mapB)

	// decoder := json.NewDecoder(bytes.NewReader(jsonWeather))
	decoder := createJsonDecoder(jsonWeather)

	// fmt.Println(decoder)

	// fmt.Println(reflect.TypeOf(decoder))

	// var weather Weather

	// err := decoder.Decode(&weather)

	// fmt.Printf("%v\n", err)

	// fmt.Printf("%+v\n", weather)

	// fmt.Println(reflect.TypeOf(weather))

	w, e := NewWeatherFromJson(decoder)

	if e != nil {
		fmt.Printf("Erro criando um novo wetaher model: %+v", e)

		return
	}

	fmt.Printf("%+v\n", w)

	fmt.Println(reflect.TypeOf(w))
}

func mockJsonWeather(mapIn map[string]interface{}) ([]byte, error) {

	json, err := json.Marshal(mapIn)

	return json, err
}

func createJsonDecoder(objJson []byte) *json.Decoder {

	return json.NewDecoder(bytes.NewReader(objJson))

}
