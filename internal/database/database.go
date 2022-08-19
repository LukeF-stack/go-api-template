package database

import (
	"database/sql"
	"example/bookAPI/internal/models/author"
	"example/bookAPI/internal/models/book"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Connectioner interface {
	Init()
}

type Connection struct {
	Db *gorm.DB
}

func (connection *Connection) Init() {
	fmt.Println("opening database connection...")
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:23306)/library")
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: db,
		}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	connection.Db = gormDB
	fmt.Println("connected to database")

	connection.Db.AutoMigrate(
		&author.Author{},
		&book.Book{},
	)
}
