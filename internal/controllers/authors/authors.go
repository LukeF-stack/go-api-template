package authors

import (
	"example/bookAPI/internal/models/author"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateAuthor(c *fiber.Ctx) error {
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var e []error
	var authorModel author.Author
	err := c.BodyParser(&authorModel)
	if err != nil {
		e = append(e, err)
	}
	statement := database.Create(&authorModel)
	if statement.Error != nil {
		e = append(e, statement.Error)
	} else {
		status = 201
	}
	c.Status(status)
	return c.JSON(types.Response{Error: e})
}

func GetAuthor(c *fiber.Ctx) error {
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var e []error
	var data types.Data = nil
	var authorModel author.Author
	query := database.First(&authorModel, c.Params("id"))
	if query.Error != nil {
		e = append(e, query.Error)
	} else {
		status = 200
	}
	data = make(map[string]any)
	data["author"] = authorModel
	c.Status(status)
	return c.JSON(types.Response{Error: e, Data: data})
}
