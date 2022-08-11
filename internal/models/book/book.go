package book

import (
	"example/bookAPI/internal/models/model_types"
)

var BookModel model_types.Table = model_types.Table{
	Name: "Book",
	Fields: []model_types.Field{
		{Column: "author", DataType: "string"},
		{Column: "genre_id", DataType: "int", FK: model_types.FK{Table: "genre", Column: "id"}},
	},
}
