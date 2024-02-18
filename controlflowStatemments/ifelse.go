package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else statements")
	loginCount := 34
	var result string
	if loginCount < 10 {
		result = "regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 count"
	}
	fmt.Println(result)

	if 89 % 2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 6; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

	// if err != nil {

	// }

}
