package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// random numbers from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random numbers from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
