package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)

	s := []interface{}{1, "James", true}
	showType(s)

}

func showType(t interface{}) {
	fmt.Printf("Type of var: %T\nValue of var: %v\n", t, t)

}

