package server

import (
	"example/bookAPI/internal/config"
	"example/bookAPI/internal/database"
	"example/bookAPI/internal/middleware"
	"example/bookAPI/internal/routes"
	"example/bookAPI/internal/server/types"
	"example/bookAPI/internal/server/utils"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	DB  *gorm.DB
	App *fiber.App
	types.Groups
}

func (server *Server) Init(connection *database.Connection) {
	server.App = fiber.New()
	server.DB = connection.Db
	server.App.Use(func(c *fiber.Ctx) error {
		utils.SetLocal[*gorm.DB](c, "db", server.DB)
		return c.Next()
	})
	// configure firebase
	firebaseAuth := config.SetupFirebase()
	server.App.Use(func(c *fiber.Ctx) error {
		utils.SetLocal[*auth.Client](c, "firebaseAuth", firebaseAuth)
		return c.Next()
	})
	server.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorisation",
	}))
	server.App.Use(middleware.AuthMiddleware)
	routes.Register(server.App, server.Groups)
	err := server.App.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
