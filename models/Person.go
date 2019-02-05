package models

type Person struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Address     *Address `json:"address"`
} 
