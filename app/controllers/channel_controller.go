package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/secrethook/backend/app/models"
	"github.com/secrethook/backend/app/queries"
	"github.com/secrethook/backend/pkg/utils"
)

func CreateNewChannel(c *fiber.Ctx) error {
	var channel models.Channel

	if err := c.BodyParser(&channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	if channel.Encryption && channel.PrivateKey == "" && channel.PublicKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "You enabled encryption but did not provide private and public keys.",
		})
	}

	channel.ID = utils.GenerateId()
	channel.CreatedAt = time.Now()
	channel.UpdatedAt = time.Now()

	validate := utils.NewValidator()

	if err := validate.Struct(channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := queries.CreateNewChannel(&channel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Channel created successfully",
		"data": channel,
	})
}

func SendMessage(c *fiber.Ctx) error {
	channelId := c.Params("channelId")

	var msg models.ChannelMessage

	msg.ChannelId = channelId
	msg.Body = c.Body()
	msg.ID = utils.GenerateId()

	validate := utils.NewValidator()

	if err := validate.Struct(msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	channelExist, err := queries.IsChannelExist(channelId)
	if err != nil && !channelExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "This channel doesn't exist",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = queries.SendNewMessage(&msg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successful create new data",
		"data": msg,
	})
}
