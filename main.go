package main

import (
	database "example/bookAPI/internal"
	"fmt"
)

func main() {
	fmt.Println("Run main!")
	db := database.GetDB()
	defer db.Close()
	fmt.Println(db.Ping())
}
