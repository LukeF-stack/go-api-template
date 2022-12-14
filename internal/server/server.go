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

// Init : initialise fiber server
func (server *Server) Init(connection *database.Connection) {
	// create fiber instance & assign server and db connection to the server struct
	server.App = fiber.New()
	server.DB = connection.Db
	// middleware to set the database as a local fiber variable (accessible through fiber context)
	server.App.Use(func(c *fiber.Ctx) error {
		utils.SetLocal[*gorm.DB](c, "db", server.DB)
		return c.Next()
	})
	// configure firebase auth
	firebaseAuth := config.SetupFirebase()
	// middleware to set the firebaseAuth struct as a local fiber variable (accessible through fiber context)
	server.App.Use(func(c *fiber.Ctx) error {
		utils.SetLocal[*auth.Client](c, "firebaseAuth", firebaseAuth)
		return c.Next()
	})
	// middleware for CORS config
	server.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorisation",
	}))
	// set auth middleware which checks for valid ID Bearer Token passed with each request
	server.App.Use(middleware.AuthMiddleware)
	// register route groups and routes
	routes.Register(server.App, server.Groups)
	// start server on port 3000
	err := server.App.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
