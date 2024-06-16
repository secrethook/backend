package queries

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/secrethook/backend/app/models"
	"github.com/secrethook/backend/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewWebhook(webhook *models.Webhook) error {
	collection := database.WebhookDatabaseClient.Collection("webhook")

	filter := bson.D{{Key: "id", Value: webhook.ID}}
	var existingWebhook models.Webhook
	var err error
	for {
		err = collection.FindOne(context.TODO(), filter).Decode(&existingWebhook)
		if err == nil {
			webhook.ID = uuid.NewString()
			continue
		} else if err == mongo.ErrNoDocuments {
			break
		} else {
			return fmt.Errorf("error checking existing webhook: %v", err)
		}
	}

	_, err = collection.InsertOne(context.TODO(), webhook)
	if err != nil {
		return fmt.Errorf("failed to insert webhook: %v", err)
	}

	log.Infof("Inserted webhook with ID: %s", webhook.ID)
	return nil
}