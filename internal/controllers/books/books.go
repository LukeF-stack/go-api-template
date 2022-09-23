package books

import (
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBook(c *fiber.Ctx) error {
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var e []string
	var bookModel book.Book
	err := c.BodyParser(&bookModel)
	if err != nil {
		e = append(e, err.Error())
	}
	statement := database.Create(&bookModel)
	if statement.Error != nil {
		e = append(e, statement.Error.Error())
	} else {
		status = 201
	}
	c.Status(status)
	return c.JSON(types.Response{Error: e})
}

func GetBook(c *fiber.Ctx) error {
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var e []string
	var data types.Data = nil
	var bookModel book.Book
	query := database.First(&bookModel, c.Query("id"))
	if query.Error != nil {
		e = append(e, query.Error.Error())
	} else {
		status = 200
	}
	data = make(map[string]any)
	if len(e) <= 0 {
		data["book"] = bookModel
	}
	c.Status(status)
	return c.JSON(types.Response{Error: e, Data: data})
}
