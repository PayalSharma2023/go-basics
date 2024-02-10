package main

import "fms"

func main () {
  a := [...]int {2,5,6}
  for i <= len(a) ; i++ {
    fmt.Println(a[i], " ")
  }
}
