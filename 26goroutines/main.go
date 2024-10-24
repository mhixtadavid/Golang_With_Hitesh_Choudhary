package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {
	// go greeter("Hello")
	// greeter("world!")
	websiteList := []string{
		"https://www.fgycvvpzsv.net",
		"https://www.hhreqe.net",
		"https://www.esliqmpg.net",
		"https://www.umpdyk.net",
		"https://www.lyneeybqrt.org",
		"https://www.techinsight.org",
		"https://www.creativeminds.io",
		"https://www.nextgenventures.com",
		"https://www.ecosolutions.net",
		"https://www.innovativedesigns.io",
	}
	for _, web := range websiteList {

		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	// for i := 0; i < 6; i++ {
// 	// 	time.Sleep(4 * time.Millisecond)
// 	// 	fmt.Println(s)

// 	// }
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOOPs in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
