// Non-Buffered Channel
// package main

// import "fmt"

// func main() {
// 	myChan := make(chan string)

// 	go func() {
// 		for {
// 			data := <-myChan
// 			fmt.Println(data)
// 		}
// 	}()

// 	myChan <- "hello world"
// 	myChan <- "I'm learning channels"
// 	myChan <- "This is extremely simple"
// }

// Basic Channel
package main

import "fmt"

func main() {
	myChan := make(chan string, 2)

	myChan <- "ABC"
	myChan <- "XYZ"

	opOne := <-myChan
	opTwo := <-myChan

	fmt.Println(opOne)
	fmt.Println(opTwo)
}

// Mutex + WaitGroup
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var mutex sync.Mutex
// var waitGroup sync.WaitGroup

// var result = []int{}

// func main() {
// 	endpoints := []string{
// 		"google.com", "twitter.com", "facebook.com",
// 	}

// 	for _, endpoint := range endpoints {
// 		go getStatusCode(endpoint, &mutex, &waitGroup)
// 		waitGroup.Add(1)
// 	}

// 	waitGroup.Wait()
// 	fmt.Println(result)
// }

// func getStatusCode(website string, mt *sync.Mutex, wg *sync.WaitGroup) {
// 	if website == "google.com" {
// 		mt.Lock()
// 		result = append(result, 200)
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("Google Fetched")
// 		mt.Unlock()
// 	} else if website == "twitter.com" {
// 		mt.Lock()
// 		result = append(result, 200)
// 		time.Sleep(time.Second * 2)
// 		fmt.Println("Twitter Fetched")
// 		mt.Unlock()
// 	} else {
// 		// For facebook -> 500 Internal Server Error
// 		mt.Lock()
// 		result = append(result, 500)
// 		time.Sleep(time.Second * 3)
// 		fmt.Println("Facebook Fetched")
// 		mt.Unlock()
// 	}

// 	wg.Done()
// }
