package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the function tutorial")

	greeter()

	// result := adder(669, 966)
	// fmt.Println("The result is:", result)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a value")

	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSpace(input)

	// result, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("Enter a valid number")
	// } else {
	// 	fmt.Printf("You entered %v\n", result)
	// }

	proSum := proAdder(2, 5, 4, 6, 8, 7, 4, 6, 5, 6, 7, 4, 2)
	fmt.Printf("the Result is: %v\n", proSum)
}

// func adder(valOne int, valTwo int) int {
// 	return valOne * valTwo
// }

func proAdder(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

func greeter() {
	fmt.Println("Hello, World!")
}
