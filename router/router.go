package router

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger" // swagger handler

	"github.com/gofiber/fiber/v2/middleware/logger"
	routes "github.com/iamatila/bigtrac/internals/routes"
)

func SetupRoutes(app *fiber.App) {

	// app.Get("/swagger/*", swagger.Handler)
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Group api calls with param '/api'
	api := app.Group("/api/v1", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	routes.SetupUserRoutes(api)
	routes.SetupTrackRoutes(api)
	routes.SetupFindRoutes(api)
}
