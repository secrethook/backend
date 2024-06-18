package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/secrethook/backend/app/controllers"
)

func ChannelRoutes(route fiber.Router) {

	route.Post("/new", controllers.CreateNewChannel)
	route.Post("/send/:channelId", controllers.SendMessage)
	
}
