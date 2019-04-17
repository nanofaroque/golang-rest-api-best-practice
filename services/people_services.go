package services

import (
	"context"
	"fmt"
	"golang-rest-api-best-practice/models"
	"golang-rest-api-best-practice/storage"
)

func GetPeople() []models.Person {
	var people []models.Person
	people = append(people, models.Person{ID: "1", Name: "Md Omar Faroque", Address: &models.Address{Street: "1305 St James Way"}})
	return people
}

func SavePerson(person *models.Person) *models.Person {
	fmt.Println(person)
	item,err :=storage.GetCollection().InsertOne(context.Background(),person)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted documents: ", item.InsertedID)
	return person
}
