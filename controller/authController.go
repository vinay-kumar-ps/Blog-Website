package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com.vinay-kumar-ps/blogbackend/database"
	"github.com.vinay-kumar-ps/blogbackend/models"
	"github.com/gofiber/fiber/v2"
)

// ValidateEmail checks if the given email is in a valid format
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return Re.MatchString(email)
}

// Register handles user registration
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	// Parse the request body
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body:", err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Invalid request body",
		})
	}

	// Check if password is less than 6 characters
	password, ok := data["password"].(string)
	if !ok || len(password) < 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Password must be at least 6 characters",
		})
	}

	// Validate email
	email, ok := data["email"].(string)
	if !ok || !ValidateEmail(strings.TrimSpace(email)) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Invalid email address",
		})
	}

	// Check if email already exists in the database
	database.DB.Where("email = ?", strings.TrimSpace(email)).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Email address already exists",
		})
	}

	// Create new user
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(email),
	}

	// Hash and set the password
	if err := user.SetPassword(password); err != nil {
		log.Println("Error hashing password:", err)
		c.Status(500)
		return c.JSON(fiber.Map{
			"Message": "Failed to hash password",
		})
	}

	// Insert user into the database
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		c.Status(500)
		return c.JSON(fiber.Map{
			"Message": "Failed to create user",
		})
	}

	// Successful response
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}
