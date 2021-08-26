package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golioth-gateway/gateway"
	"github.com/golioth-gateway/goliothMongo"
)

var db goliothMongo.Repository

func init() {

	time.Sleep(3 * time.Second)

	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	telemetryCollection := os.Getenv("TELEMETRY_COLLECTION")

	dbPort, _ := strconv.Atoi(dbPortStr)

	db = goliothMongo.NewMongoClient(dbHost, dbPort)

	if err := db.Connect(); err != nil {
		fmt.Println("***[ERRO] : Could not initialize mongodbapi client:")
		return
	}

	db.DropCollection(dbName, telemetryCollection)

	fmt.Println("db: ", db)

}

func main() {

	portEnv := os.Getenv("GATEWAY_PORT")

	port, err := strconv.Atoi(portEnv)

	if err != nil {

		fmt.Printf("Error trying to get the Gateway PORT Env Variable. Value got = %s. %v\n", portEnv, err)

		return
	}

	gtw, err := gateway.NewGateway(port, db)

	if err != nil {
		fmt.Printf("Error instantiating a new Gateway: %v\n", err)

		return
	}

	if err := gtw.Run(); err != nil {

		fmt.Printf("Error trying to run the Gateway: %v\n", err)

		return
	}
}
