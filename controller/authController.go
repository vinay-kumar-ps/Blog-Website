package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com.vinay-kumar-ps/blogbackend/database"
	"github.com.vinay-kumar-ps/blogbackend/models"
	"github.com.vinay-kumar-ps/blogbackend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
			"message": "Invalid request body",
		})
	}

	// Check if password is less than 6 characters
	password, ok := data["password"].(string)
	if !ok || len(password) < 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be at least 6 characters",
		})
	}

	// Validate email
	email, ok := data["email"].(string)
	if !ok || !ValidateEmail(strings.TrimSpace(email)) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email address",
		})
	}

	// Check if email already exists in the database
	database.DB.Where("email = ?", strings.TrimSpace(email)).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email address already exists",
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
			"message": "Failed to hash password",
		})
	}

	// Insert user into the database
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	// Successful response
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// Parse the request body
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body:", err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address Doesnt Exist,Kindly create an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect  Password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "You have suucessfully loggined",
		"user":    user,
	})
}
	type Claims struct{
		jwt.StandardClaims
	}
