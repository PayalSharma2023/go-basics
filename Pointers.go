package main

import "fmt"

type Counter struct {
  hits int 
}

func IncreaseCounter(counter*Counter) {
  counter.hits := Counter.hits + 1
  fmt.Println(counter.hits)
}

func main() {
  counterOne = Counter{}
  IncreaseCounter(&counterOne)
  fmt.Println(counterOne.hits)
}
