package main

func main() {
	a := 1
	b := 2

	if a < b {
		println(b)
	} else {
		println(a)
	}

	switch a {
	case 1:
		println(1)
	case 2:
		println(2)
	case 3:
		println(3)
	default:
		println("do not exists")
	}
}
