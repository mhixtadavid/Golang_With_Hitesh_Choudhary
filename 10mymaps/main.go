package main

import "fmt"

func main() {
	fmt.Println("Welcome to Map tutorials")

	languages := make(map[string]string)

	languages["CM"] = "Chemistry"
	languages["PY"] = "Physics"
	languages["BIO"] = "Biology"

	fmt.Println(languages)
	fmt.Println(languages["BIO"])

	delete(languages, "BIO")
	fmt.Println(languages)

	// peep into how loops operate

	for key, value := range languages {
		fmt.Printf("the key is %v and the value is %v\n", key, value)
	}

	languages = append(languages, map[string], "BIO""Biology")

}
