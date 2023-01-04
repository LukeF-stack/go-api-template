package books

import (
	"example/bookAPI/internal/models/book"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	logger "example/bookAPI/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateBook : creates a new book entry in the db using the passed req data
func CreateBook(c *fiber.Ctx) error {
	// get db connection from fiber context
	database := utils.GetLocal[*gorm.DB](c, "db")
	// set default response status and error slice
	var status = 500
	var e []string
	// create book struct using the book model
	var bookModel book.Book
	// parse the request fields into the book struct (& is used to point to the above struct)
	err := c.BodyParser(&bookModel)
	if err != nil {
		e = append(e, err.Error())
	}
	// write the new book struct to the database
	statement := database.Create(&bookModel)
	if statement.Error != nil {
		// append sql error to error slice
		e = append(e, statement.Error.Error())
	} else {
		// set http status to 201 (new entity created)
		status = 201
	}
	// set the fiber status
	c.Status(status)
	// send JSON response using the base "types.Response" struct
	return c.JSON(types.Response{Error: e})
}

func GetBook(c *fiber.Ctx) error {
	storeLogger := utils.GetLocal[*logger.Logger](c, "firestoreClient")
	storeLogger.Log("from the book route", utils.GetLocal[string](c, "uuid"))
	database := utils.GetLocal[*gorm.DB](c, "db")
	var status = 500
	var e []string
	var data types.Data = nil
	var bookModel book.Book
	// preload is used to load the author struct associated with the book
	// .First gets the first book that matches the query params
	// c.Query is used to access the id query passed with the request
	query := database.Preload("Author").First(&bookModel, c.Query("id"))
	if query.Error != nil {
		e = append(e, query.Error.Error())
	} else {
		status = 200
	}
	// make a map for the response data
	data = make(map[string]any)
	// send the book back in the response if no error occurred
	if len(e) <= 0 {
		data["book"] = bookModel
	}
	c.Status(status)
	return c.JSON(types.Response{Error: e, Data: data})
}
