package main

import fmt

func greet(name string) string {
	return "Hello, " + name
}

func welcome(name string) string {
  return "welcome, " + " to the programming world " + name 
}

func main() {
	message := greet("Payal Sharma")
  welcome := welcome("PAYAL")
	fmt.Println(message)
}
