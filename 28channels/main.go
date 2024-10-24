package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go lang - KDAdev.ng")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	//receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, IsChanelOpen := <-myCh

		fmt.Println(IsChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {

		myCh <- 0
		close(myCh)
		// myCh <- 6

		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
