package main

import (
	"fmt"
)

// there is no inheritance in golang : no super or parent
type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func main() {
	Henry := User{"Henry", 23, true, "henry454@gmail.com"}
	fmt.Println(Henry)
	fmt.Printf("Henry's details are : %v \n", Henry)
	fmt.Printf("Name is %v , Age is %v", Henry.Name, Henry.Age)

}
