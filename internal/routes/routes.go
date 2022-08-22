package routes

import (
	"example/bookAPI/internal/controllers/books"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/book/create", books.CreateBook)
}
