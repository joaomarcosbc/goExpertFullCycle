package main

import "fmt"

func main() {

	doubleSum := func() int {
		return sum(1,5,6,7,8,10) * 2
	}()

	fmt.Println(sum(1,5,6,7,8,10))
	fmt.Println(doubleSum)
}

func sum(numbers ...int) int{
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}