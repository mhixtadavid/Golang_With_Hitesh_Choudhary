package main

import "fmt"

func main() {
	fmt.Println("welcome to strut tutorial")

	kaydee := User{"David", "david@go.com", true, 78}
	fmt.Println(kaydee)
	fmt.Printf("The details of david is: %v\n", kaydee)
	fmt.Printf("The details of david is: %+v\n", kaydee)
	fmt.Printf("The name is %v and the age is %v\n", kaydee.Name, kaydee.Age)
	kaydee.GetStatus()
	kaydee.NewMail()
	fmt.Printf("The name is %v and the email is %v\n", kaydee.Name, kaydee.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is useractive: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "kjhk@hghg.com"
	fmt.Println("Thee new email is: ", u.Email)
}
