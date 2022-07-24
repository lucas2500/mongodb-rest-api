package models

import (
	"fmt"
	"mongodb-rest-api/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertDocument(document interface{}, filter interface{}, CollectionName string) bool {

	// Get database connection
	var client = database.Client

	db := client.Database("ServiceA").Collection(CollectionName)
	opts := options.Update().SetUpsert(true)

	// Close database connection
	// defer database.DisconnectFromMongo(client)

	_, err := db.UpdateOne(database.DbCtx, filter, document, opts)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DeleteDocument(filter interface{}, CollectionName string) (bool, int64) {

	// Get database connection
	var client = database.Client

	db := client.Database("ServiceA").Collection(CollectionName)

	rows, err := db.DeleteOne(database.DbCtx, filter)

	if err != nil {
		fmt.Println(err)
		return false, 0
	}

	return true, rows.DeletedCount
}

func FindDocument(CollectionName string) (bool, []bson.M) {

	// Get database connection
	var client = database.Client
	var results []bson.M

	db := client.Database("ServiceA").Collection(CollectionName)

	cursor, err := db.Find(database.DbCtx, bson.D{{}})

	if err != nil {
		fmt.Println(err)
		return false, results
	}

	if err = cursor.All(database.DbCtx, &results); err != nil {
		fmt.Println(err)
		return false, results
	}

	return true, results
}
