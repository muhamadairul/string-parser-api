package routes

import "github.com/gofiber/fiber/v2"

// PublicRoutes registers public routes (no auth required).
func PublicRoutes(a *fiber.App) {
	a.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
