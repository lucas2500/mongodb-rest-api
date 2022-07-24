package models

import (
	"fmt"
	"mongodb-rest-api/database"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertDocument(document interface{}, filter interface{}, DocumentName string) bool {

	// Get database connection
	var client = database.Client

	db := client.Database("ServiceA").Collection(DocumentName)
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

func DeleteDocument(filter interface{}, DocumentName string) (bool, int64) {

	// Get database connection
	var client = database.Client

	db := client.Database("ServiceA").Collection(DocumentName)

	rows, err := db.DeleteOne(database.DbCtx, filter)

	if err != nil {
		fmt.Println(err)
		return false, 0
	}

	return true, rows.DeletedCount
}
