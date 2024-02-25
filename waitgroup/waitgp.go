package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex  //pointer
func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endPoint string) {
	defer wg.Done()
	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", endPoint, err)
		return
		// fmt.Println("OOPS in endpoint")
	} 

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %v\n", endPoint, err)
		return
	}

	// mut.Lock()
	// res = append(res, endPoint)
	// mut.Unlock()
	mut.Lock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)
	mut.Unlock()
	
	//%d for integer and %s for string

}
