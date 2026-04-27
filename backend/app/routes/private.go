package routes

import (
	"github.com/gofiber/fiber/v2"
	stringparser "github.com/muhamadairul/string-parser-api/app/controllers/string-parser"
)

// PrivateRoutes registers API routes for the string parser.
func PrivateRoutes(a *fiber.App) {
	api := a.Group("/api")

	api.Post("/parse", stringparser.Parse)
	api.Get("/history", stringparser.History)
}
