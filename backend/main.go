package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/muhamadairul/string-parser-api/app/middleware"
	"github.com/muhamadairul/string-parser-api/app/models/entities"
	"github.com/muhamadairul/string-parser-api/app/routes"
	"github.com/muhamadairul/string-parser-api/app/utils/db"
	"github.com/muhamadairul/string-parser-api/app/utils/parser"
	"github.com/muhamadairul/string-parser-api/app/utils/server"
	"github.com/muhamadairul/string-parser-api/configs"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	if _, err := db.OpenDBConnection(); err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	db.EnableGlobalTimestamps(db.Query)

	if err := db.Query.AutoMigrate(&entities.ParsedResult{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	log.Println("Database schema migrated successfully")

	// Load provincial capitals from runtime configuration file
	capitals, err := parser.LoadCapitals("configs/capitals.json")
	if err != nil {
		log.Fatalf("Failed to load capitals config: %v", err)
	}
	log.Printf("Loaded %d provincial capitals from config", len(capitals))

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app, capitals)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		server.StartServer(app)
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}
