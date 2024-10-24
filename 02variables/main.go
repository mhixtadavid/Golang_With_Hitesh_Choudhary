package main

import "fmt"

const LoginToken string = "HelloDav" // public => this can be called anywhere

func main() {
	var username string = "Kingdavid"
	fmt.Println(username)
	fmt.Printf("Variable is a type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is a type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is a type: %T \n", smallVal)

	var smallFloat float64 = 255.45545545555455455455455455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is a type: %T \n", smallFloat)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is a type: %T \n", anotherVar)

	// implicit type

	var website = "david.king"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is a type: %T \n", LoginToken)
}
