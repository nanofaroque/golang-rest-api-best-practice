package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api-best-practice/controllers"
	_ "golang-rest-api-best-practice/storage"
	"log"
	"net/http"
)

func main() {
	/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection:=storage.GetCollection()
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err!=nil {

	}
	id := res.InsertedID
	fmt.Println("I am printing",id)*/

	router := mux.NewRouter()
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/people",controllers.SavePerson).Methods("POST")
	fmt.Print("Server has started")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Print("Server has started")
}

