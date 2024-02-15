package main

import "fmt"

type MyStruct struct {
	x int
	y int
}

type coordinate struct {
	X, Y int
}

//simple function
func modifyNormal(x int, y int, MyStruct *MyStruct) {
	MyStruct.x += x
	MyStruct.y += y
}

//Reciever function (Pointer)

func (MyStruct *MyStruct) modifyWithRec(x int, y int) {
	MyStruct.x += x
	MyStruct.y += y
}

// Reciever Function (value)

func (start coordinate) calculateDistance(end coordinate) coordinate {
	return coordinate{start.X - end.X, start.Y - end.Y}
}

func main() {
	myStruct := MyStruct{x: 10, y: 10}

	modifyNormal(1, 2, &myStruct)

	myStruct.modifyWithRec(3, 4)
	fmt.Println(myStruct)

	startCord := coordinate{14, 2}
	endCord := coordinate{5, 1}

	distance := startCord.calculateDistance(endCord)
	fmt.Println(distance)
}
