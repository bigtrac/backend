package userHandler

import (
	// "bytes"
	// "encoding/json"

	// "encoding/json"

	// "net/http"

	// "net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/iamatila/bigtrac/database"
	model "github.com/iamatila/bigtrac/internals/models"
	// mail "github.com/iamatila/bigtrac/internals/handlers/mailer"
)

// UpdateLocation func create a find
// @Description Create a Find
// @Tags Find
// @Accept json
// @Produce json
// @Success 200 {object} model.Find
// @router /api/v1/find/update/:device_id [patch]
func UpdateLocation(c *fiber.Ctx) error {
	db := database.DB
	device_id := c.Params("device_id")
	var find model.Find

	// Find by device_id
	if err := db.Where("device_id = ?", device_id).First(&find).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Find not found",
		})
	}
	// // Find the device by device_id
	// result := db.Where("device_id = ?", deviceID).First(&device)
	// if result.Error != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"error": "Device not found",
	// 	})
	// }

	// Check if the device is active
	if !find.Active {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Device is not active",
		})
	}

	// Save the updated find
	db.Save(&find)

	// Return the created find
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Find found",
		// "message": "Email confirmation OTP Sent to your email",
		// "data":    find,
	})

}
