package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2) //buffered channel
	wg := &sync.WaitGroup{}

	//if somebody is listening then only I will take value else I will give deadlock
	// 	fmt.Println(<-myCh)
	// 	myCh <- 5
	// 	fmt.Println(<-myCh)
	wg.Add(2)

	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh) //returns zero
		//myCh <- 5
		//myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
