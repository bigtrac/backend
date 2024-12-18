package routes

import (
	// "log"
	// "os"

	"github.com/gofiber/fiber/v2"

	// jwtware "github.com/gofiber/jwt/v3"
	// jwtware "github.com/gofiber/contrib/jwt"
	findHandler "github.com/iamatila/bigtrac/internals/handlers"
	// "github.com/joho/godotenv"
)

func SetupFindRoutes(router fiber.Router) {
	find := router.Group("/find")
	// Create a User
	find.Patch("/update/:device_id", findHandler.UpdateLocation)
}
