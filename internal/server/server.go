package server

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/routes/books"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	DB  *gorm.DB
	App *fiber.App
}

func (server *Server) Init(connection *database.Connection) {
	server.App = fiber.New()
	server.DB = connection.Db
	books.Init(server.App, server.DB)
	err := server.App.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
