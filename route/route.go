package route

import (
	"github.com.vinay-kumar-ps/blogbackend/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register",controller.Register)
}