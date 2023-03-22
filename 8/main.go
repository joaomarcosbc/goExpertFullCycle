package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(50, 10)
	if  err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func sum(a, b int) (int, error) {
	
	if a + b >= 50 {
		return a + b, errors.New("Sum is greater than 50")
	}
	return a + b, nil
}