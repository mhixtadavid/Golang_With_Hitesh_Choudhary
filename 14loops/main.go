package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loop classes")

	// days := []string{"Sunday", "Tuesday", "Thursday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("The index is %v and the value is %v\n", index, day)
	// }

	roguevalue := 1
	for roguevalue < 10 {

		if roguevalue == 3 {
			goto Hello
		}

		if roguevalue == 5 {
			roguevalue++
			// break
			continue
		}
		fmt.Println(roguevalue)
		roguevalue++
	}
Hello:
	fmt.Println("you got caught")

}
