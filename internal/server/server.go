package server

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/routes/books"
	"github.com/gofiber/fiber/v2"
)

func Init(connection *database.Connection) {
	app := fiber.New()
	books.Init(app, connection)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
