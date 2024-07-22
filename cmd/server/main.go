package handler

import (
	"net/http"
	// "scanNstore/cmd/routes"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "GET, POST",
		AllowCredentials: false,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Route("/auth", func(router fiber.Router) {
	// 			routes.SetupAuthRoutes(router)
	// 		})

	// Convert the Fiber app to an HTTP handler and serve the request
	adaptor.FiberApp(app)(w, r)
}
