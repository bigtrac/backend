package routes

import (
	// "log"
	// "os"

	"github.com/gofiber/fiber/v2"

	// jwtware "github.com/gofiber/jwt/v3"
	// jwtware "github.com/gofiber/contrib/jwt"
	trackHandler "github.com/iamatila/bigtrac/internals/handlers"
	// "github.com/joho/godotenv"
)

func SetupTrackRoutes(router fiber.Router) {
	track := router.Group("/track")
	// Create a Track
	track.Post("/log", trackHandler.CreateTrack)
}
