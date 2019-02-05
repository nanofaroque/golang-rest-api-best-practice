package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	people = append(people, models.Person{ID: "1", Name: "Md Omar Faroque", Address:&models.Address{Street: "1305 St James Way"}})
	json.NewEncoder(w).Encode(people)
}