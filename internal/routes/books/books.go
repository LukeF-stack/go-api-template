package books

import (
	"encoding/json"
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/server/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, database *gorm.DB) {
	app.Get("/book/create", func(c *fiber.Ctx) error {
		success := false
		var error []error
		newBook := book.Book{Name: "New Book", AuthorID: 1}
		result := database.Create(&newBook)
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
		fmt.Println("testing")
		return c.SendString("About")
	})
}
