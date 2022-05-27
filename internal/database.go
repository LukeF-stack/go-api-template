package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDB() *sql.DB {
	fmt.Println("opening database connection...")
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/book")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
