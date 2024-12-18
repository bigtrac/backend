package userHandler

import (
	// "bytes"
	// "encoding/json"

	// "encoding/json"

	// "net/http"

	// "net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"

	"github.com/iamatila/bigtrac/database"
	model "github.com/iamatila/bigtrac/internals/models"
	// mail "github.com/iamatila/bigtrac/internals/handlers/mailer"
)

// CreateUser func create a user
// @Description Create a User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @router /api/v1/users/register [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	find := new(model.Find)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}

	// Add a uuid to the user
	user.Userid = uuid.New().String()
	user.Active = true

	// Create the User and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	// finding track save
	// Store the body in the user and return error if encountered
	errk := c.BodyParser(find)
	if errk != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}

	find.Userid = user.Userid
	find.DeviceId = user.DeviceId
	find.Name = user.Name
	find.Active = user.Active

	// Create the finding track and return error if encountered
	err = db.Create(&find).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create find",
			"data":    err,
		})
	}

	// Return the created user
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Account created successfully",
		// "message": "Email confirmation OTP Sent to your email",
		// "data":    user,
	})

}
