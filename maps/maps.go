package main

import (
	"fmt"
)

func main() {
	languages := make(map[string]string) //key value pair
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list : ", languages)
	fmt.Println("JS : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loop iteration and use of range in loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key, value is %v \n", value)
	}
}
