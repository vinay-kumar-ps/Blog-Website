package route

import (
	// "github.com.vinay-kumar-ps/blogbackend/middleware"
	"github.com/gofiber/fiber/v2"
 controller	"github.com.vinay-kumar-ps/blogbackend/controller"

)

// SetUp defines the routes for the application
func SetUp(app *fiber.App) {
	// Public routes
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// Protected routes, requires authentication
	// protected := app.Group("/", middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
}
