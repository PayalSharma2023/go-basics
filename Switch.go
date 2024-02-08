package main

import "fmt"

func Area(length int , breadth int) int {
  return length * breadth 
}
func main () {
  switch area := Area(20, 54) {
    case area <= 300 : 
        fmt.Println("Area = ", area)
    case area >= 400 :
        fmt.Println("Area appropriate for work = ", area)
    
  }
}
