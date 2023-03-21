package main

import "fmt"

func main() {

	salarys := map[string]int{"James": 4800, "John": 3200, "JM": 2400}
	// salarys := map[string]int{}
	// salarys := make(map[string]int)

	delete(salarys, "JM")

	salarys["Joao"] = 5000

	for name, salary := range salarys {
		fmt.Printf("Name: %s, salary: %d\n", name, salary)
	}

	for _, salary := range salarys {
		fmt.Printf("Salary: %d\n", salary)
	}

}