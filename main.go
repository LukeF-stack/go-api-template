package main

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/server"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Main: Starting Application")
	connection := new(database.Connection)
	failure := make(chan error)
	go connection.Init(failure)
	fmt.Println("Main: Waiting for db connection to finish")
	err := <-failure
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Main: Database Connected")
	apiServer := new(server.Server)
	apiServer.Init(connection)
}
