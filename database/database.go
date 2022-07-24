package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DbCtx  context.Context
)

func InitDatabase() {

	var err error
	DbCtx = context.Background()

	Client, err = mongo.Connect(DbCtx, options.Client().ApplyURI("mongodb+srv://lucas2500:followyou@cluster0.7k8ipyc.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("There was an error when tryting to connect to MongoDB!!")
	}

	fmt.Println("Connection established successfully!!")
}

func DisconnectFromMongo() {

	defer func() {
		if err := Client.Disconnect(DbCtx); err != nil {
			log.Fatal("There was an error when trying to disconnect from database!! ", err)
		}
	}()
}
