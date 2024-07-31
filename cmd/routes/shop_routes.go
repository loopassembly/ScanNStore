package routes

import (
	"github.com/gofiber/fiber/v2"
	"scanNstore/cmd/controllers"
)

func SetupAuthRoutes(router fiber.Router) {
	

	router.Post("/receipts", controllers.CreateReceipt)
	router.Get("/receipts/:id", controllers.GetReceipt)
	router.Get("/receipts", controllers.ListReceipts)
	router.Post("/upload", controllers.UploadImage)

}
