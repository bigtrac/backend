package main

import (
	"log"
	"os"

	// "strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iamatila/bigtrac/database"
	"github.com/iamatila/bigtrac/router"
	"github.com/joho/godotenv"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// allowedOrigins := strings.Join([]string{"https://admin.landtify.com", "http://localhost:9000"}, ",")
	// allowedOrigins := strings.Join([]string{"https://admin.landtify.com", "http://localhost:9000"}, ",")

	// // Configure CORS middleware with allowed origins
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{allowedOrigins},
	// }))
	// Initialize default config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		// AllowOrigins: []string{"https://admin.landtify.com", "http://localhost:9000"},
		// AllowOrigins: []string{allowedOrigins},
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	// Listen on PORT 3000
	// app.Listen(port)
	app.Listen(`0.0.0.0:` + port)
}
