package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	//var no_one int = 2
	//var no_two float64 = 4.5

	// fmt.Println("the sum of number one and number two is : ", no_one + no_two)

	//random number

	// rand.Seed(time.Now().UnixMicro()) //for truly random number
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNumber)

}
