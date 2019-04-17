package main

import (
	_ "./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"golang-rest-api-best-practice/controllers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	fmt.Print("Server has started")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Print("Server has started")
}

// Display all from the people var

