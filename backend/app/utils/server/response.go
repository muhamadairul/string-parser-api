package server

import (
	"github.com/gofiber/fiber/v2"
)

// Response is the standard API response structure.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseSuccess sends a 200 OK JSON response.
func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	if message == "" {
		message = "Berhasil mendapatkan data!"
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ResponseBadRequest sends a 400 Bad Request JSON response.
func ResponseBadRequest(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Request tidak valid!"
	}
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  "error",
		Message: message,
	})
}

// ResponseError sends a 500 Internal Server Error JSON response.
func ResponseError(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Maaf, terjadi kesalahan server!"
	}
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  "error",
		Message: message,
	})
}

// ResponseNotFound sends a 404 Not Found JSON response.
func ResponseNotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Data tidak ditemukan!"
	}
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:  "error",
		Message: message,
	})
}
