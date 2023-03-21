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
	fmt.Printf("Type of e variable: %T\n", e)
	fmt.Printf("Value of e variable: %v\n", e)
	fmt.Printf("Type of f variable: %T\n", f)
}
