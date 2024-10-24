package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Wekcome to the if-else class")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "Welcome"
	}
	fmt.Println(result)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Enter a valid number")
	}

	if int(number)%2 == 0 {
		fmt.Printf("%v is even\n", number)
	} else {
		fmt.Printf("%v is odd\n", number)
	}

	if number < 10 {
		fmt.Printf("%v is less than 10\n", number)
	} else {
		fmt.Printf("%v is greater than 10\n", number)
	}
}
