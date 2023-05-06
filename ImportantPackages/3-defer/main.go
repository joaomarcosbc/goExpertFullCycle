package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.joaomarcos.dev")
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
}
