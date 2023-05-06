package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("myfile.md")
	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello World")
	size, err := f.Write([]byte("Writing data on file"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created. Size: %d bytes\n", size)

	f.Close()

	// Reading Files
	file, err := os.ReadFile("myfile.md")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Reading parts

	file2, err := os.Open("myfile.md")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}


	err = os.Remove("myfile.md")
	if err != nil {
		panic(err)
	}
}
