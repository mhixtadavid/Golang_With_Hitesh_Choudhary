package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to this JSON course")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	kdacourses := []course{
		{"React JS Bootcamp", 299, "Kaydeedevelopers.ng", "kaydee123", []string{"Web-dev", "JS"}},
		{"MERN JS Bootcamp", 345, "Kaydeedevelopers.ng", "dev321", []string{"full=stack", "JS"}},
		{"Angular JS Bootcamp", 199, "Kaydeedevelopers.ng", "kay123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(kdacourses, "kda", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		 {
             "coursename": "React JS Bootcamp",
             "Price": 299,
             "website": "Kaydeedevelopers.ng", 
             "tags": ["Web-dev","JS"]
     }
	`)

	var kdaCourse course

	checkValid := json.Valid((jsonDataFromWeb))
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &kdaCourse)
		fmt.Printf("%#v\n", kdaCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the Value is %v and the type is %T\n", k, v, v)
	}
}
