package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const myurl string = "https://www.geeksforgeeks.org/golang/"

func main() {
	fmt.Println("Handling web requests with http")
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T\n", response) //response has a header, body
	defer response.Body.Close()
	//databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//content := string(databyte)
	//fmt.Println(content)

	fmt.Println(myurl)
	//Parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["pagename"])
	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Payal",
	}
	anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl)

}
