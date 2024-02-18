package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("We will learn to get input from user")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of your experience in college")

	//similar to try catch we have comma ok , err ok, err err here
	//also known as valrus operator
	input, _ := reader.ReadString('\n') // we use _ if we want to ignore error
	fmt.Println("Thank for rating : ", input)
	fmt.Printf("type of rating is : %T \n", input) // here if we want to add 1 in input we cannot as type of input is string do it so for this we have another topc of conversion

	//numRating := input + 1 // error
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating ",numRating + 1)
	}

}
