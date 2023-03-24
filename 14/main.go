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

type Animal struct {
	Name 	string
	Age 	int
}

type LivinBeing interface {
	changeAge()
}

func (a Animal) changeAge() {
	a.Age = 4
	fmt.Println(a)
}

func (u User) changeAge() {
	u.Age = 21
	fmt.Println(u)
}

func changingAge(livingBeing LivinBeing) {
	livingBeing.changeAge()
}

func main() {
	user := User{
		Name: "James",
		Age: 20,
		Status: true,
	}

	dog := Animal{
		Name: "Chica",
		Age: 3,
	}

	user.Name = "John"
	user.Adress.Street = "Urbano Neto"
	user.City = "Aracaju"

	changingAge(dog)
	changingAge(user)

	fmt.Println(user)
	fmt.Println(dog)
}