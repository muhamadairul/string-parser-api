package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // Auto-load .env

	"github.com/muhamadairul/string-parser-api/app/middleware"
	"github.com/muhamadairul/string-parser-api/app/models/entities"
	"github.com/muhamadairul/string-parser-api/app/routes"
	"github.com/muhamadairul/string-parser-api/app/utils/db"
	"github.com/muhamadairul/string-parser-api/app/utils/server"
	"github.com/muhamadairul/string-parser-api/configs"
)

func main() {
	// Initialize Fiber with config
	config := configs.FiberConfig()
	app := fiber.New(config)

	// Open database connection
	if _, err := db.OpenDBConnection(); err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	db.EnableGlobalTimestamps(db.Query)

	// Auto-migrate schema
	if err := db.Query.AutoMigrate(&entities.ParsedResult{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	log.Println("Database schema migrated successfully")

	// Register middlewares
	middleware.FiberMiddleware(app)

	// Register routes
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server
	if os.Getenv("STAGE_STATUS") == "dev" {
		server.StartServer(app)
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}
