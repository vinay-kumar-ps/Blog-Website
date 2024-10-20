package controller

import (
	"fmt"

	"github.com.vinay-kumar-ps/blogbackend/database"
	"github.com.vinay-kumar-ps/blogbackend/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse the body")
	}
	if err :=database.DB.Create(&blogpost).Error;err!=nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invalid payload",
		})
	}
		return c.JSON(fiber.Map{
			"message":"Congratulation! Your Post Is live",
		})
	}

