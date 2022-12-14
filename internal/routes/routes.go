package routes

import (
	"example/bookAPI/internal/controllers/authors"
	"example/bookAPI/internal/controllers/books"
	"example/bookAPI/internal/server/types"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, groups types.Groups) {
	// create route groups
	Author := app.Group("/author/")
	Book := app.Group("/book/")
	// create a group map
	groups = make(map[string]fiber.Router)
	// assign the groups to the group map
	groups["Author"] = Author
	groups["Book"] = Book
	// declare routes and corresponding controller functions
	groups["Book"].Post("/create", books.CreateBook)
	groups["Book"].Get("/get", books.GetBook)
	groups["Author"].Post("/create", authors.CreateAuthor)
	groups["Author"].Get("/get", authors.GetAuthor)
}
