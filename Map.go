package main

import "fmt"

func main() {
  Student := map [string] int {
    "Leonardo" : 112325, 
    "Darwin" : 536372,
    "Priya" : 73839,
  }

  for key, value := range Student {
    fmt.println(key, " " , value)
  }

  fmt.println()
  snacks := make(map[int]string)
  snacks[1] := "Chips",
  snacks[2] := "Creamrole"

  delete(snacks, 2)

  for key, value := range snacks {
    fmt.println(key, " ", value)
  }

  _, found := Student[112325]
  if !found {
    fmt.println("No student found")
  } else {
    fmt.println("student found")
  }
}
