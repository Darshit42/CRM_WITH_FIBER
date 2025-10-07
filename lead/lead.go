package lead

import (
	"github.com/Darshit42/CRM_WITH_FIBER/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db.Create(&lead).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create lead"})
	}

	return c.Status(201).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	if err := db.First(&lead, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lead not found"})
	}

	db.Delete(&lead)
	return c.JSON(fiber.Map{"message": "Lead successfully deleted"})
}
