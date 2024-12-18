package userHandler

import (
	// "bytes"
	// "encoding/json"

	// "encoding/json"

	// "net/http"

	// "net/http"

	"github.com/gofiber/fiber/v2"

	// "github.com/google/uuid"

	"github.com/iamatila/bigtrac/database"
	model "github.com/iamatila/bigtrac/internals/models"
	// mail "github.com/iamatila/bigtrac/internals/handlers/mailer"
)

// CreateTrack func create a track
// @Description Create a Track
// @Tags Track
// @Accept json
// @Produce json
// @Success 200 {object} model.Track
// @router /api/v1/track/log [post]
func CreateTrack(c *fiber.Ctx) error {
	db := database.DB
	track := new(model.Track)

	// Store the body in the track and return error if encountered
	err := c.BodyParser(track)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}

	// // Add a uuid to the track
	// user.Userid = uuid.New().String()
	// user.Active = true

	// Create the track and return error if encountered
	err = db.Create(&track).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create track",
			"data":    err,
		})
	}

	// mail.UserOtpMail(c, user.Email, OTP)

	// Return the created track
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Location created successfully",
		// "message": "Email confirmation OTP Sent to your email",
		// "data":    track,
	})

}
