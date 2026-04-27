package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// StartServer starts Fiber without graceful shutdown (dev mode).
func StartServer(a *fiber.App) {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server starting on %s", addr)
	if err := a.Listen(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// StartServerWithGracefulShutdown starts Fiber with graceful shutdown (prod mode).
func StartServerWithGracefulShutdown(a *fiber.App) {
	// For simplicity, same as StartServer but can be extended with os.Signal handling
	StartServer(a)
}
