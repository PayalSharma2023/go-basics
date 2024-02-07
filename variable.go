package main

import "fmt"

func main() {
	var fName string = "Payal"
	var lName = "Sharma"

	age := 19

	fmt.Println("Hello ", fName, " ", lName, ". You are ", age, " years old")

	// ok Error Idiom
	partOne, other := 10, 20
	fmt.Println("Part One ", partOne, " ", " Other ", other)

	partTwo, other := 30, 40
	fmt.Println("Part Two ", partTwo, " ", " Other ", other)

	// Multiple Variables

	var (
		userName = "psyuio"
		userId   = "176768989"
	)

	fmt.Println("Username ", userName)
	fmt.Println("UserID ", userId)

	// Ignoring values

	a, b, _, d := 'A', 'B', 'C', 'D'

	fmt.Println(string(a))
	fmt.Println(string(b))
	

}
