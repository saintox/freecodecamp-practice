package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/saintox/go-practice/freecodecamp/crm-fiber-server/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBConnection
	var leads []Lead

	db.Find(&leads)

	ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConnection

	var lead Lead
	db.Find(&lead, id)

	if lead.Name == "" {
		ctx.Status(500).Send("lead not found")
		return
	}

	ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) {
	db := database.DBConnection
	lead := new(Lead)

	if err := ctx.BodyParser(lead); err != nil {
		ctx.Status(503).Send(err)
		return
	}

	db.Create(&lead)

	ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConnection

	var lead Lead

	db.First(&lead, id)

	if lead.Name == "" {
		ctx.Status(500).Send("lead not found with id " + ctx.Params("id"))
		return
	}

	db.Delete(&lead)

	ctx.Send("lead deleted successfully")
}
