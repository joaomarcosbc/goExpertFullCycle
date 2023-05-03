package main

import "fmt"

const a = "Hello World"

type ID int // Create Types

var (
	b bool 		= true 
	c int  		= 10 
	d string 	= "Joao" 
	e float64 	= 1.2 
	f ID 		= 1
)
 
func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v:= range myArray {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}
}
