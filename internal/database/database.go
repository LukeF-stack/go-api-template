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

func (connection *Connection) Init(error chan<- error) {
	fmt.Println("opening database connection...")
	// open base sql driver connection
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		// send the error back through the channel
		error <- err
	}
	db.SetMaxIdleConns(10)
	// open a new GORM connection with the sql driver connection
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}), &gorm.Config{})
	if err != nil {
		error <- err
	}
	// set the connection db to the gorm instance
	connection.Db = gormDB
	fmt.Println("connected to database")
	// send the error back as nil through the channel
	error <- nil
}
