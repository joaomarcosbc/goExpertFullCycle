package main

import "fmt"

func main() {

	salaries := map[string]int{"James": 4800, "John": 3200, "JM": 2400}
	// salaries := map[string]int{}
	// salaries := make(map[string]int)

	delete(salaries, "JM")

	salarys["Joao"] = 5000

	for name, salary := range salaries {
		fmt.Printf("Name: %s, salary: %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("Salary: %d\n", salary)
	}

}