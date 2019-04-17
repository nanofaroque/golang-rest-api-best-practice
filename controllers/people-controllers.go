package controllers

import (
	"encoding/json"
	"golang-rest-api-best-practice/services"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPeople())

}