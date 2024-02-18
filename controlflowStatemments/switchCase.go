package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //for some random value
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is : ",diceNumber)

	switch diceNumber{
	    case 1 :
			fmt.Println("Dice value is one you can open")
		case 2 :
			fmt.Println("You can move 2 steps")
			fallthrough
		case 3 :
			fmt.Println("You can move 3 steps")
			fallthrough
		case 4 :
			fmt.Println("You can move 4 steps")
		case 5 :
			fmt.Println("You can move 5 steps")
        case 6 :
			fmt.Println("You can move 6 steps")

	}
}
