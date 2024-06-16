package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/secrethook/backend/app/controllers"
)

func WebhookRoutes(route fiber.Router) {
	route.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, POST request received")
	})
	route.Post("/new", controllers.CreateNewWebhook)
}
