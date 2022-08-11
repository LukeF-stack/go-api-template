package main

import (
	"example/bookAPI/internal/database"
	"fmt"
)

func main() {
	fmt.Println("Run main!")
	connection := new(database.Connection)
	connection.Init()
	defer connection.Unmount()
}
