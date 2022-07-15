package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Connectioner interface {
	Init()
}

type Connection struct {
	Db *sql.DB
}

func (connection *Connection) Init() {
	fmt.Println("opening database connection...")
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	} else {
		connection.Db = db
		fmt.Println("connected to database")
	}
}

func (connection *Connection) Unmount() {
	connection.Db.Close()
}
