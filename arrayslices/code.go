package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "lichi"
	//fruitList[4] = "strawberry"  error
	fmt.Println("fruit List : ", fruitList)
	fmt.Println(len(fruitList))

	var vegetableList = []string{"Potato", "Peas", "Ladyfinger"}
	fmt.Println(vegetableList)

	fmt.Printf("type of fruitList is :%T \n ", fruitList)
	vegetableList = append(vegetableList, "capcicum") // here's how we can add more items in a array
	fmt.Println(vegetableList)

	vegetableList = append(vegetableList[1:])
	fmt.Println(vegetableList)

	HighScore := make([]int, 4)
	HighScore[0] = 45
	HighScore[1] = 780
	HighScore[2] = 90
	HighScore[3] = 134
	//HighScore[4] = 787 error
	HighScore = append(HighScore, 676)

	fmt.Println(sort.IntsAreSorted(HighScore))

	sort.Ints(HighScore) //only aplicable in slice

	fmt.Println(HighScore)
	fmt.Println(sort.IntsAreSorted(HighScore))
}
