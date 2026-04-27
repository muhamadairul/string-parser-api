package stringparser

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhamadairul/string-parser-api/app/models/entities"
	"github.com/muhamadairul/string-parser-api/app/utils/db"
	"github.com/muhamadairul/string-parser-api/app/utils/parser"
	"github.com/muhamadairul/string-parser-api/app/utils/server"
)

// ParseRequest defines the expected JSON request body.
type ParseRequest struct {
	Input string `json:"input"`
}

// Parse handles POST /api/parse
// Takes a raw input string, parses it into Name/Age/City, stores to DB, returns result.
func Parse(c *fiber.Ctx) error {
	var req ParseRequest
	if err := c.BodyParser(&req); err != nil {
		return server.ResponseBadRequest(c, "Format JSON tidak valid")
	}

	if strings.TrimSpace(req.Input) == "" {
		return server.ResponseBadRequest(c, "Input tidak boleh kosong")
	}

	// Parse: right-to-left, no regex, no replace, max 5 vars
	rawName, rawAge, rawCity := parser.Parse(strings.ToUpper(req.Input))

	// Enrich city with province if it's a provincial capital
	enrichedCity := parser.EnrichCity(rawCity)

	// Apply fixed-width formatting:
	// Name: 30 char left-aligned, Age: 3 char right-aligned, City: 20 char left-aligned
	fmtName := fmt.Sprintf("%-30s", rawName)
	fmtAge := fmt.Sprintf("%3s", rawAge)
	fmtCity := fmt.Sprintf("%-20s", enrichedCity)

	// Truncate if exceeds CHAR limit (safety)
	if len(fmtName) > 30 {
		fmtName = fmtName[:30]
	}
	if len(fmtAge) > 3 {
		fmtAge = fmtAge[:3]
	}
	if len(fmtCity) > 20 {
		fmtCity = fmtCity[:20]
	}

	// Persist to database
	record := entities.ParsedResult{
		Name: fmtName,
		Age:  fmtAge,
		City: fmtCity,
	}
	if err := db.Query.Create(&record).Error; err != nil {
		return server.ResponseError(c, "Gagal menyimpan data: "+err.Error())
	}

	return server.ResponseSuccess(c, fiber.Map{
		"id":   record.ID,
		"name": record.Name,
		"age":  record.Age,
		"city": record.City,
		"raw": fiber.Map{
			"name": rawName,
			"age":  rawAge,
			"city": enrichedCity,
		},
		"created_at": record.CreatedAt,
	}, "Parsing berhasil!")
}
