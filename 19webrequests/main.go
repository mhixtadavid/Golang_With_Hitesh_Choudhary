package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("Welcome to the course on web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	// databytes, err := ioutil.ReadAll(response.Body) // deprecated
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
