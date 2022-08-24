package main

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Main: Starting Application")
	connection := new(database.Connection)
	awaitConn := make(chan bool)
	err := make(chan error)
	go connection.Init(awaitConn, err)
	fmt.Println(<-err)
	fmt.Println("Main: Waiting for db connection to finish")
	<-awaitConn
	fmt.Println("Main: Database Connected")
	apiServer := new(server.Server)
	apiServer.Init(connection)
}
