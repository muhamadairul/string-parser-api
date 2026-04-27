package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadairul/string-parser-api/app/utils/server"
)

// NotFoundRoute registers 404 route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		return server.ResponseNotFound(c, "Route tidak ditemukan")
	})
}
