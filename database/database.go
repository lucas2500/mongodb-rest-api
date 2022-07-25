package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	DbCtx    context.Context
	MongoURI string
)

func InitDatabase() {

	var err error
	DbCtx = context.Background()

	Client, err = mongo.Connect(DbCtx, options.Client().ApplyURI(MongoURI))

	if err != nil {
		log.Fatal("There was an error when tryting to connect to MongoDB!!")
	}

	fmt.Println("Connection established successfully!!")
}

func DisconnectFromMongo() {

	defer func() {
		if err := Client.Disconnect(DbCtx); err != nil {
			log.Fatal("There was an error when trying to disconnect from MongoDB!! ", err)
		}
	}()
}
