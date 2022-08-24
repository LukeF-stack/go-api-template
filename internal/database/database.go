package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connectioner interface {
	Init()
}

type Connection struct {
	Db *gorm.DB
}

func (connection *Connection) Init(finished chan<- bool, errors chan<- error) {
	fmt.Println("opening database connection...")
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library")
	if err != nil {
		errors <- err
	}
	db.SetMaxIdleConns(10)
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}), &gorm.Config{})
	if err != nil {
		errors <- err
	}
	connection.Db = gormDB
	fmt.Println("connected to database")
	finished <- true
}
