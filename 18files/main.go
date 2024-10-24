package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files handling tutorials")

	content := "This needs to go into the file here"

	file, err := os.Create("./mygofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("Length id: ", length)
	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	// databyte, err := ioutil.ReadFile(filename) //depricated method
	databyte, err := os.ReadFile(filename) // current method
	checkNilErr(err)

	fmt.Println("Text data inside the file is:- ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
