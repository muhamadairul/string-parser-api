package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// FiberMiddleware registers all global Fiber middlewares.
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// CORS — allow all origins (dev mode); tighten in production
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
			AllowHeaders: "Origin,Content-Type,Accept,Authorization",
		}),

		// Request logger
		logger.New(),

		// Panic recovery
		recover.New(),
	)
}
