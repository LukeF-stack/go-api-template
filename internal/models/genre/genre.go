package genre

import (
	"example/bookAPI/internal/models/model_types"
)

var GenreModel model_types.Table = model_types.Table{
	Name: "Genre",
	Fields: []model_types.Field{
		{Column: "title", DataType: "string"},
	},
}
