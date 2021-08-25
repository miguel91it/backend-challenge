package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golioth-gateway/gateway"
)

func main() {

	portEnv := os.Getenv("PORT")

	port, err := strconv.Atoi(portEnv)

	if err != nil {

		fmt.Printf("Error trying to get the Gateway PORT Env Variable. Value got = %s. %v\n", portEnv, err)

		return
	}

	dbPort := os.Getenv("DB_PORT")

	fmt.Printf("PORT: %d\nDB_PORT: %s\n", port, dbPort)

	gtw, err := gateway.NewGateway(port)

	if err != nil {
		fmt.Printf("Error instantiating a new Gateway: %v\n", err)

		return
	}

	if err := gtw.Run(); err != nil {

		fmt.Printf("Error trying to run the Gateway: %v\n", err)

		return
	}
}
