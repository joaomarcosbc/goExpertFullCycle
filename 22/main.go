package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"one", "three", "five"}

	for i, v := range numbers {
		println(i, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Infinity")
	}
}
