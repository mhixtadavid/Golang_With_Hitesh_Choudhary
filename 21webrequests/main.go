package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content length is", response.ContentLength)

	var responseCount strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseCount.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseCount.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"Coursename":"Learning post in Golang",
			"Price": 0,
			"Platform": "Kaydeedev.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "david")
	data.Add("lastname", "oshin")
	data.Add("email", "david@oshin.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
