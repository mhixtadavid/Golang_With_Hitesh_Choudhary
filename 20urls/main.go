package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://hjhk.com:4000/learn?coursename=reactjs&paymentid=sodf786287468"

func main() {
	fmt.Println("Url handling class")

	fmt.Println(myurl)

	// parsing

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of the query params are: %T\n", qparams)

	fmt.Println(qparams["Coursename"])

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "david.com",
		Path:    "/youtube videos/upth/ago iwoye",
		RawPath: "user=kingdavid",
	}

	fmt.Println(partsOfUrl.String())
}
