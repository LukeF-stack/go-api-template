package types

import "github.com/gofiber/fiber/v2"

type Response struct {
	Error []string
	Data  Data
}

type Groups map[string]fiber.Router

type Data map[string]any
