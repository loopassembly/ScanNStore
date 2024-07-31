package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"scanNstore/cmd/initializers"
	"scanNstore/cmd/routes"
)

func init() {

	initializers.ConnectDB()
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", 
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	// Mount the microservice routes under "/api"
	micro := app.Group("/api")

	// Setup auth routes
	routes.SetupAuthRoutes(micro)

	// Health check route
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	// Index route
	// Index route
app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Welcome to ScanNStore")
})


	// Catch-all route for 404 errors
	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exist on this server", path),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
