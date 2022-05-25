package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/saintox/go-practice/freecodecamp/crm-fiber-server/database"
	"github.com/saintox/go-practice/freecodecamp/crm-fiber-server/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect DB")
	}

	fmt.Println("Database connected")
	database.DBConnection.AutoMigrate(&lead.Lead{})
	fmt.Println("Migration success")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(8080)

	defer database.DBConnection.Close()
}
