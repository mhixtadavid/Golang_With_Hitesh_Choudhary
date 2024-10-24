package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your score here: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Bravo! Your score is -", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Added +1 to your score -", numRating+1)
	}

}
