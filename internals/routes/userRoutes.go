package routes

import (
	// "log"
	// "os"

	"github.com/gofiber/fiber/v2"

	// jwtware "github.com/gofiber/jwt/v3"
	// jwtware "github.com/gofiber/contrib/jwt"
	userHandler "github.com/iamatila/bigtrac/internals/handlers"
	// "github.com/joho/godotenv"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/users")
	// Create a User
	user.Post("/register", userHandler.CreateUser)
}
