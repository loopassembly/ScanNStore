package controllers

import (
	"scanNstore/cmd/initializers"
	"scanNstore/cmd/models"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreateReceipt(c *fiber.Ctx) error {
	var payload models.Receipt


	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Validate the payload
	// errors := models.ValidateStruct(&payload)
	// if errors != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	// }

	// Set timestamps
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()


	result := initializers.DB.Create(&payload)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": payload})
}

// GetReceipt retrieves a receipt by its ID along with its items
func GetReceipt(c *fiber.Ctx) error {
	id := c.Params("id")
	var receipt models.Receipt

	// Find the receipt by ID and preload its items
	if err := initializers.DB.Preload("Items").First(&receipt, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Receipt not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": receipt})
}




func ListReceipts(c *fiber.Ctx) error {
	var receipts []models.Receipt

	// Find all receipts without preloading items
	if err := initializers.DB.Find(&receipts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not retrieve receipts"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": receipts})
}

