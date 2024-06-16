package database

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/secrethook/backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var WebhookDatabaseClient *mongo.Database

func ConnectToMongodb() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongodbConnectionURL, _ := utils.ConnectionURLBuilder("mongodb")
	opts := options.Client().ApplyURI(mongodbConnectionURL).SetServerAPIOptions(serverAPI)
	MongodbClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Errorf("failed connected to MongoDB!", "err", err)
	}
	if err := MongodbClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Errorf("failed connected to MongoDB!", "err", err)
	}

	WebhookDatabaseClient = MongodbClient.Database("webhook")

	log.Info("You successfully connected to MongoDB!")
}
