package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/services"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPeople())
}