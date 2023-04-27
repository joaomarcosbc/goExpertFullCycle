package main

import "fmt"

type Account struct {
	balance float64
}

// kind of constructor function
// question: whats the difference between initialize with pointers and not, today!!!
func NewAccount() *Account {
	return &Account{balance: 0}
}


////////////////////////////////////////////////
func (a Account) simulate(value float64) {
	a.balance += value
	fmt.Printf("Should be %v\n", a.balance)
}

func (a *Account) add(value float64) float64{
	a.balance += value
	return a.balance
}

func main() {
	account := NewAccount()

	account.add(12)
	account.simulate(300)

	fmt.Println(account.balance)


}