package books

import (
	"encoding/json"
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBook(c *fiber.Ctx) error {
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var error []error
	newBook := book.Book{Name: "New Book", AuthorID: 1}
	result := database.Create(&newBook)
	if result.Error != nil {
		error = append(error, result.Error)
		return result.Error
	} else {
		status = 201
	}
	response, err := json.Marshal(types.Response{Error: error})
	if err != nil {
		return err
	}
	c.Status(status)
	return c.Send(response)
}
