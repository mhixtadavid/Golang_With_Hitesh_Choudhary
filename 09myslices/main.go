package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the slices tutorial")

	var bookList = []string{"Chemistry", "Physics", "Biology"}
	fmt.Printf("the data type is: %T\n", bookList)

	bookList = append(bookList, "Economics", "Geography")
	fmt.Println(bookList)

	bookList = append(bookList[:3])
	fmt.Println(bookList)

	highScore := make([]int, 4)

	highScore[0] = 576
	highScore[1] = 536
	highScore[2] = 465
	highScore[3] = 322
	// highScore[4] = 999

	highScore = append(highScore, 862, 676, 878)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	// how to remove a value from slices based on index

	var courses = []string{"Chem", "Phy", "Bio", "Maths", "Eng"}
	fmt.Println(courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
