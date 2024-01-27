package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishavqwerty7/BookwormApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://bookworm:bookworm123@cluster0.d8bnhnd.mongodb.net/?retryWrites=true&w=majority"
var dbName = "bookworm"
var colName = "readlist"

var collection *mongo.Collection

func init() {

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("connection is ready")

}

// helper
func insertOneBookHelper(book model.Bookworm) {

	inserted, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted book with id:", inserted.InsertedID)
}

func updateOneBookHelper(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"read": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated count id", updated.ModifiedCount)

}

func deleteOneBookHelper(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete count id", deleteCount)

}

func deleteManyHelper() {
	filter := bson.D{{}}
	deletedResult, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted count is", deletedResult.DeletedCount)
}

func getAllBooksHelper() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var Books []primitive.M

	for curr.Next(context.Background()) {
		var book bson.M

		err := curr.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		Books = append(Books, book)
	}

	return Books
}

//Actual controller

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var book model.Bookworm

	_ = json.NewDecoder(r.Body).Decode(&book)

	insertOneBookHelper(book)

	json.NewEncoder(w).Encode(book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneBookHelper(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-methods", "DElETE")

	params := mux.Vars(r)
	deleteOneBookHelper(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	allMovies := getAllBooksHelper()

	json.NewEncoder(w).Encode(allMovies)

}

func DeleteAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteManyHelper()
}
