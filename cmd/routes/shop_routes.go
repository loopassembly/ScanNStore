// auth.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"scanNstore/cmd/controllers"

	
)
// all auth routes including oauth
func SetupAuthRoutes(router fiber.Router) {
	router.Get("/status", controllers.Trail)


}