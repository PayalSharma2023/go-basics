package main

import "fmt"

const ClassId = 1234

var SchoolName string = "Kendriya Vidyalaya No.1 "

func main() {
	fmt.Println("Welcome to the world of Programming in Go language")

	var Name string = "Payal Sharma"
	fmt.Println("Hi ", Name)
	fmt.Printf("Name is of type : %T \n", Name)

	var isStudyingInSchool bool = false
	fmt.Println(isStudyingInSchool)
	fmt.Printf("isStudyingInShool is of type : %T \n", isStudyingInSchool)

	var smallint uint8 = 255 //return only small values will give error at 256
	fmt.Println(smallint)

	var float float32 = 56.672145479
	fmt.Println(float)//returns only presion upto 5 decimal places for more precision use float64

	fmt.Println("School name :", SchoolName)

	//default value and some aliases
	var anotherShool string
	fmt.Println(anotherShool) // is returning blank
	fmt.Printf("anotherSchool is of type : %T \n", anotherShool) // is giving string as type

	//implicit type in declaring variables
	var feature = "painting" //here lexer will automatically analyze that feature is of datatype string 
	//we can use way of declaring variables globally
	fmt.Println(feature)
	fmt.Printf("feature is of type : %T \n", feature)

	//no var style in writing code
	Books := 5
	fmt.Println("no. of books : ", Books)//this is aplicable only inside the method 
}
