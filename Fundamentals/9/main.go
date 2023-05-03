package main

import "fmt"

func main() {
	fmt.Println(sum(1,5,6,7,8,10))
}

func sum(numbers ...int) int{
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}