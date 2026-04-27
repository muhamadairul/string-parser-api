package stringparser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadairul/string-parser-api/app/models/entities"
	"github.com/muhamadairul/string-parser-api/app/utils/db"
	"github.com/muhamadairul/string-parser-api/app/utils/server"
)

// History handles GET /api/history.
func History(c *fiber.Ctx) error {
	var results []entities.ParsedResult
	if err := db.Query.Order("created_at DESC").Limit(20).Find(&results).Error; err != nil {
		return server.ResponseError(c, "Gagal mengambil data history: "+err.Error())
	}
	return server.ResponseSuccess(c, results, "Berhasil mendapatkan history!")
}
