package main

import "fmt"

func main() {
	fmt.Println("welcome to strut tutorial")

	kaydee := User{"David", "david@go.com", true, 78}
	fmt.Println(kaydee)
	fmt.Printf("The details of david is: %v\n", kaydee)
	fmt.Printf("The details of david is: %+v\n", kaydee)
	fmt.Printf("The name is %v and the age is %v\n", kaydee.Name, kaydee.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
