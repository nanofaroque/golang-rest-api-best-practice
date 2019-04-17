package main

import (
	_ "./controllers"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-rest-api-best-practice/controllers"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err!= nil{
		return
	}
	fmt.Print(client)
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Println(id)

	router := mux.NewRouter()
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	fmt.Print("Server has started")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Print("Server has started")
}

// Display all from the people var
