package main

import "fmt"

type User struct {
	Name string
	Age int
	Status bool
}

func main() {
	user := User{
		Name: "James",
		Age: 20,
		Status: true,
	}

	user.Name = "John"

	fmt.Println(user.Name)
}