package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch and case tutorial")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	// fmt.Print(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("You have %v\n", diceNumber)
	case 2:
		fmt.Printf("You have %v\n", diceNumber)
	case 3:
		fmt.Printf("You have %v\n", diceNumber)
	case 4:
		fmt.Printf("You have %v\n", diceNumber)
	case 5:
		fmt.Printf("You have %v\n", diceNumber)
	case 6:
		fmt.Printf("You have %v\n", diceNumber)

	}
}
