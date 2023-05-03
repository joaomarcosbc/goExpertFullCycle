package main

import (
	"first/math"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := math.Sum(10, 20)

	fmt.Println(uuid.New())

	println(s)
}
