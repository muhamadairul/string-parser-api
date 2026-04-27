package routes

import (
	"github.com/gofiber/fiber/v2"
	stringparser "github.com/muhamadairul/string-parser-api/app/controllers/string-parser"
)

// PrivateRoutes registers API routes.
func PrivateRoutes(a *fiber.App, capitals map[string]string) {
	api := a.Group("/api")

	api.Post("/parse", stringparser.ParseHandler(capitals))
	api.Get("/history", stringparser.History)
}
