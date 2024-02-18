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

	//to remove value from slice
	HighScore = append(HighScore[:3], HighScore[4:]...)
	fmt.Println(HighScore)
	//based on index
	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // ... is symbol of variadics
	fmt.Println(courses)

	days := []string{"Mon", "Tue", "Wed", "Fri", "Thurs"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for _, i := range days{
		fmt.Println(i)
	}

	for index, days := range days{
		fmt.Printf("index is %v and day is %v \n", index, days)
	}
}
