package main

import "fmt"

func main() {

	// Memory -> Adress -> Value

	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	*b = 30
	*&a = 40
	

	fmt.Println(pointer, b, *b, &a, a)

}