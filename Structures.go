package main

import "fmt"

type Student struct{
  roll_no int
  name   string 
}
func main () {
  student1 := Student{11256, "Payal"}
  fmt.Println(student1)
}
