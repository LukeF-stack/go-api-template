package main

import (
	"database/sql"
	"example/bookAPI/internal/models/author"
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/models/job"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
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
	fmt.Println("connected to database")
	fmt.Println("migrating...")
	err = gormDB.AutoMigrate(
		&author.Author{},
		&book.Book{},
		&job.Job{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migration complete")
}
