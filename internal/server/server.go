package server

import (
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/routes"
	"example/bookAPI/internal/server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	DB  *gorm.DB
	App *fiber.App
}

func (server *Server) Init(connection *database.Connection) {
	server.App = fiber.New()
	server.DB = connection.Db
	server.App.Use(func(c *fiber.Ctx) error {
		utils.SetLocal[*gorm.DB](c, "db", server.DB)
		return c.Next()
	})
	routes.Register(server.App)
	err := server.App.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
