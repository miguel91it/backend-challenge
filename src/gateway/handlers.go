package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	j, _ := json.Marshal(weather)

	var weatherAsMap map[string]interface{}

	json.Unmarshal(j, &weatherAsMap)

	// db.CreateDocuments(dbName, telemetryCollection, []map[string]interface{}{
	db.CreateDocuments("golioth-challenge", "weatherTelemetry", []map[string]interface{}{

		weatherAsMap,
	})

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Weather Data received succesfuly")
}

func GetWeatherByFilter(w http.ResponseWriter, r *http.Request) {

	var body struct {
		DeviceId  string `json:"device_id"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	// fmt.Printf("Body: %+v\n", body)

	dateFilter := make(bson.M)

	startTime, err := parseStringToDate(body.StartDate)

	if err != nil {
		// fmt.Println(err)

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to parse the start_date filter to a valid date object: %s", err.Error())

		return

	} else {

		dateFilter["$gte"] = startTime.UTC().Unix()
	}

	endTime, err := parseStringToDate(body.EndDate)

	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to parse the end_date filter to a valid date object: %s", err.Error())

		return

	} else {

		dateFilter["$lte"] = endTime.UTC().Unix()
	}

	// fmt.Println("\nFilter: ", dateFilter, "\n")

	weathers, _ := db.GetDocsByFilter("golioth-challenge", "weatherTelemetry", bson.M{
		"timestamp": dateFilter,
	})

	// j, _ := json.Marshal(weathers)
	// fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(weathers)

}

func parseStringToDate(dateStr string) (time.Time, error) {

	// only for running outside container, to get the host timezone
	// zone, _ := time.Now().Local().Zone()

	// only for running inside container because container has the UTC timezone
	zone := "-00"

	dateStr += zone + ":00"

	dateParsedTime, err := time.Parse(time.RFC3339, dateStr)

	if err != nil {

		return time.Now(), fmt.Errorf("error trying to parse string date to date object: %v", err)
	}

	return dateParsedTime, nil

}
