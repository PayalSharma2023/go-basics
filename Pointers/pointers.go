package main

import "fmt"

func main() {
	fmt.Println("Pointers are used to make changes or refer to the original value instead of making a copy")
	var ptr *int
	fmt.Print(ptr) // will show <nil> as output

	myNum := 56
	var Ptr = &myNum //& is used for reference

	fmt.Println("value of actual pointer : ", Ptr) // Pointers actual value is memory address of that value
	fmt.Println("value inside the pointer : ", *Ptr) //* with poitnter will give the value that is inside in the pointer

	*Ptr = *Ptr * 4 //we can make changes in the value of actual variable
	fmt.Println("new value is : ", myNum)
}
