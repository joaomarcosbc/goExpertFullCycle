package main

type MyNumber int

//constraint
type Number interface{
	~int | ~float64
}

//func sum[T int | float64]
func sum[T Number] (m map[string]T) T {	
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// can only be used for equality
func compare[T comparable] (a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Juba": 7000, "Joao": 4500, "Lau": 5000}
	m2 := map[string]float64{"Juba": 700.0, "Joao": 450.0, "Lau": 500.0}
	m3 := map[string]MyNumber{"Juba": 7000, "Joao": 4500, "Lau": 5000}

	println(sum(m), sum(m2), sum(m3))
	println(compare(10, 10))
	println(compare("Test", "James"))

	// s := []any == s := []interface{}{1, "James", true}

}