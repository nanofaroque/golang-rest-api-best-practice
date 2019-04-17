package controllers

import (
	"encoding/json"
	"golang-rest-api-best-practice/models"
	"golang-rest-api-best-practice/services"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetPeople())
}

func SavePeople(w http.ResponseWriter, r *http.Request) {
	address:=models.Address{"1305 St jayson"}
	person:=models.Person{"slkdjfslkdjfsd","faroque",&address}
	services.SavePerson(&person)
}