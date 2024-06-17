package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/secrethook/backend/app/models"
	"github.com/secrethook/backend/app/queries"
	"github.com/secrethook/backend/pkg/utils"
)

func CreateNewWebhook(c *fiber.Ctx)  error {
	var webhook models.Webhook

	if err := c.BodyParser(&webhook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": "Invalid request body",
		})
	}

	if webhook.Encryption && webhook.PrivateKey == "" && webhook.PublicKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": "You enabled encryption but did not provide private and public keys.",
		})
	}
	
	webhook.ID = uuid.NewString()
	webhook.CreatedAt = time.Now()
	webhook.UpdatedAt = time.Now()

	validate := utils.NewValidator()

	if err := validate.Struct(webhook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := queries.CreateNewWebhook(&webhook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Webhook created successfully",
		"webhook": webhook,
	})
}