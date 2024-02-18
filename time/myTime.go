package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("we will learn how we can work with time")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2024, time.December, 2, 3, 4, 4, 43, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Tuesday"))
}
