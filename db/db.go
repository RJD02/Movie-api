package db

import (
	"context"
	"fmt"
	"log"

	"github.com/RJD02/mongoapi/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var configParams = config.GetConfig()
var connectionString = "mongodb+srv://admin-" + configParams["MONGODB_USER"] + ":" + configParams["MONGODB_PASSWORD"] + "@cluster0.lkxsz.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const collName = "watchlist"

func GetCollection() *mongo.Collection {
	var collection *mongo.Collection

	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongo db
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	collection = dbClient.Database(dbName).Collection(collName)

	fmt.Println("Collection instance is ready")

	return collection
}
