package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome on board to Go language", input)
}
