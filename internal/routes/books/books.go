package books

import (
	"encoding/json"
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/server/types"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App, connection *database.Connection) {
	app.Get("/book/create", func(c *fiber.Ctx) error {
		success := false
		var error []error
		newBook := book.Book{Name: "New Book", AuthorID: 1}
		result := connection.Db.Create(&newBook)
		if result.Error != nil {
			error = append(error, result.Error)
			panic(result.Error)
		} else {
			success = true
		}
		response, err := json.Marshal(types.Response{Error: error, Success: success})
		if err != nil {
			panic(err)
		}
		return c.Send(response)
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About")
	})
}
