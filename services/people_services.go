package services

import "golang-rest-api-best-practice/models"

func GetPeople() []models.Person  {
	var people []models.Person
	people = append(people, models.Person{ID: "1", Name: "Md Omar Faroque", Address:&models.Address{Street: "1305 St James Way"}})
    return people

}
