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
	Fields []Field
}

type Field struct {
	Column   string
	DataType string
}

func (database *Database) Init() {
	var fieldFields []Field
	fieldFields = append(fieldFields, Field{Column: "field", DataType: "string"})
	var book Table = Table{Fields: fieldFields}
	for i := 0; i < 2; i++ {
		database.Tables = append(database.Tables, book)
	}
	for i := 0; i < len(database.Tables); i++ {
		jsonObj, err := json.Marshal(database.Tables[i])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(jsonObj))
		}
	}
}
