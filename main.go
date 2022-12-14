package main

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/server"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Main: Starting Application")
	// instantiate a db struct instance
	connection := new(database.Connection)
	// create channel to send error back from async db connection
	failure := make(chan error)
	// asynchronous db connection using go routine and GORM
	go connection.Init(failure)
	fmt.Println("Main: Waiting for db connection to finish")
	// assign the error channel once message comes back
	err := <-failure
	// if there is a db connection error log a fatal error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Main: Database Connected")
	// instantiate a server struct instance
	apiServer := new(server.Server)
	// start the server and pass the db connection
	apiServer.Init(connection)
}
