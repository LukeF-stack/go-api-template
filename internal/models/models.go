package models

import (
	"encoding/json"
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/models/model_types"
	"fmt"
	"log"
)

type Constraint interface {
	model_types.Table
}

type Iterator interface {
	forEach(callback callback)
}

type Database struct {
	Tables []model_types.Table
}

type callback func(model_types.Table, int)

func (database *Database) forEach(callback callback) {
	for i := 0; i < len(database.Tables); i++ {
		callback(database.Tables[i], i)
	}
}

func (database *Database) Init() {
	database.Tables = append(database.Tables,
		book.BookModel,
	)
	database.forEach(func(table model_types.Table, i int) {
		jsonObj, err := json.Marshal(table)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(jsonObj))
		}
	})
}
