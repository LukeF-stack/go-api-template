package main

import (
	database "example/bookAPI/internal"
	"fmt"
)

func main() {
	fmt.Println("Run main!")
	connection := new(database.Connection)
	connection.Init()
	defer connection.Unmount()

}
