package models

import (
	"encoding/json"
	"fmt"
	"log"
)

type Database struct {
	Tables []Table
}

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Column   string
	DataType string
}

var ModelTables []Table

var BookModel = Table{Name: "Book", Fields: []Field{{Column: "author", DataType: "string"}}}

func (database *Database) Init() {
	database.Tables = ModelTables
	for i := 0; i < len(database.Tables); i++ {
		jsonObj, err := json.Marshal(database.Tables[i])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(jsonObj))
		}
	}
}
