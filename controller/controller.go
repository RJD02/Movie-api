package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RJD02/mongoapi/config"
	"github.com/RJD02/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var configParams = config.GetConfig()

var connectionString = "mongodb+srv://admin-" + configParams["MONGODB_USER"] + ":" + configParams["MONGODB_PASSWORD"] + "@cluster0.lkxsz.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const collName = "watchlist"

// Most important
var collection *mongo.Collection

// connect with mongodb

// runs only one time
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongo db
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	collection = dbClient.Database(dbName).Collection(collName)

	fmt.Println("Collection instance is ready")

}

// mongodb helpers - file

// insert 1 record
func insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a movie with id:", inserted.InsertedID)
}

// update a record
func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count: ", result)

}

// delete all records
func deleteAllMovies() {

	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delete count:", deleteResult.DeletedCount)
}

// get all records
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// actual controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	deleteAllMovies()
	json.NewEncoder(w).Encode("Deleted everything")
}
