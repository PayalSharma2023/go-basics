package main

import "fmt"

func main() {
  // Correct the map declaration using equals signs instead of colons
  Student := map[string]int{
    "Leonardo":  112325,
    "Darwin":   536372,
    "Priya":    73839,
  }

  // Use the correct Println function from the fmt package
  for key, value := range Student {
    fmt.Println(key, " ", value)
  }

  fmt.Println()

  // Create a new map for snacks and use equals signs for assignment
  snacks := make(map[int]string)
  snacks[1] = "Chips"
  snacks[2] = "Creamrole"

  // Delete the snack with key  2
  delete(snacks,  2)

  // Iterate over the snacks map and print its contents
  for key, value := range snacks {
    fmt.Println(key, " ", value)
  }

  // Correctly check for a student by using a string key
  _, found := Student["Leonardo"]
  if !found {
    fmt.Println("No student found")
  } else {
    fmt.Println("student found")
  }
}
