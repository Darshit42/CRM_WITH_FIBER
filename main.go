package main

import (
	"fmt"

	"github.com/Darshit42/CRM_WITH_FIBER/database"
	"github.com/Darshit42/CRM_WITH_FIBER/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to conect to db")
	}
	fmt.Println("Connected to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println(("Database Migrated"))
}

func main() {
	app := fiber.New()
	initDB()
	setupRoutes(app)
	app.Listen(":3000")

	defer database.DBConn.Close()
}
