package main

import "fmt"

type Adress struct {
	Street 	string
	Number 	string
	City 	string
	State 	string
}

type User struct {
	Name 	string
	Age 	int
	Status 	bool
	Adress 	
}

func main() {
	user := User{
		Name: "James",
		Age: 20,
		Status: true,
	}

	user.Name = "John"
	user.Adress.Street = "Urbano Neto"
	user.City = "Aracaju"

	fmt.Println(user)
}