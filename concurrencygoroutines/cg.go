package main

import (
	"fmt"
	"time"
)
//Concurrency Vs parallelism
func main() {//go routine 
	go greeter("hello") //it is firing up different threads
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond) //using this to recieve those threads
		fmt.Println(s)
	}
}
