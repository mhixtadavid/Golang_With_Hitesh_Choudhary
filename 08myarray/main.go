package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array lessons")

	// you have to specify the number you want to pass into the array eg[4]

	var bookList [4]string

	bookList[0] = "Chemistry"
	bookList[2] = "Physics"
	bookList[3] = "Biology"

	fmt.Println("The list of books are:", bookList)
	fmt.Println("The list of books are:", len(bookList))

	var FoodList = [10]string{"rice", "beans", "garri"}
	fmt.Println("Food list include:", len(FoodList))
}
