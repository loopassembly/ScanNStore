package controllers

import (
	
	"github.com/gofiber/fiber/v2"
)



func Trail(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"aessage": "Welcome to Golang, Fiber, and GORM",
	})
}
