package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://www.geeksforgeeks.org/golang/"

func main() {
	fmt.Println("Handling web requests with http")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T\n", response) //response has a header, body
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

}
