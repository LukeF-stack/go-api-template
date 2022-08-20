package main

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Main: Starting Application")
	awaitConn := make(chan bool)
	connection := new(database.Connection)
	go connection.Init(awaitConn)
	fmt.Println("Main: Waiting for db connection to finish")
	<-awaitConn
	fmt.Println("Main: Database Connected")
	server.Init(connection)
}
