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

func CreateNewChannel(channel *models.Channel) error {
	var err error
	collection := database.ChannelDatabaseClient.Collection("channel")

	filter := bson.D{{Key: "id", Value: channel.ID}}
	var existingChannel models.Channel
	for {
		err = collection.FindOne(context.TODO(), filter).Decode(&existingChannel)
		if err == nil {
			channel.ID = uuid.NewString()
			continue
		} else if err == mongo.ErrNoDocuments {
			break
		} else {
			log.Errorf("error checking existing channel: %v", err)
			return fmt.Errorf("error checking existing channel: %v", err)
		}
	}

	_, err = collection.InsertOne(context.TODO(), channel)
	if err != nil {
		log.Errorf("failed to insert channel: %v", err)
		return fmt.Errorf("failed to insert channel: %v", err)
	}

	log.Infof("Inserted channel with ID: %s", channel.ID)
	return nil
}

func SendNewMessage(msg *models.ChannelMessage) error {
	var err error

	messageCollection := database.ChannelDatabaseClient.Collection("message")

	_, err = messageCollection.InsertOne(context.TODO(), msg)
	if err != nil {
		log.Errorf("failed to insert message: %v", err)
		return fmt.Errorf("failed to insert message: %v", err)
	}

	return nil
}

func IsChannelExist(channelId string) (bool, error) {

	var err error

	channelCollection := database.ChannelDatabaseClient.Collection("channel")

	filter := bson.D{{Key: "id", Value: channelId}}

	var ChannelData models.ChannelMessage
	err = channelCollection.FindOne(context.TODO(), filter).Decode(&ChannelData)
	if err == mongo.ErrNoDocuments {
		return false, fmt.Errorf("this channel doesn't exist")
	} else if err != nil {
		log.Errorf("error checking existing channel: %v", err)
		return true, fmt.Errorf("error checking existing channel: %v", err)
	}

	return true, nil
}
