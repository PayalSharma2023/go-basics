// package main

// import (
// 	"fmt"
// 	"sync"
// )
// var wg sync.WaitGroup
// type People interface {
// 	GetName() string
// 	ChangePassword(NewPass string)
// }

// // func (p *People) GetName() {

// // fmt.Printf("Name is : %s\n", p )
// // }
// type Student struct {
// 	Name     string
// 	Email    string
// 	Password string
// 	Rollno   int
// }

// func (S Student) GetName() string {
// 	return fmt.Sprintln("Student Name is : ", S.Name)
// }

// func (S Student) ChangePassword(NewPass string) {
// 	S.Password = NewPass
// 	fmt.Println("New Password is : ", S.Password)
// }

// type Teacher struct {
// 	Name     string
// 	Email    string
// 	Password string
// }

// func (T Teacher) GetName() string {
// 	return fmt.Sprintln("Teacher Name is : ", T.Name)

// }

// func (T Teacher) ChangePassword(NewPass string) {
// 	T.Password = NewPass
// 	fmt.Println("Changed password is : ", T.Password)
// }

// func main() {
// 	//created one instance
// 	Payal := Student{
// 		Name:     "Payal Sharma",
// 		Email:    "payalsharma@gmail.com",
// 		Password: "6778bsu",
// 		Rollno:   112327,
// 	}

// 	Shama := Teacher{
// 		Name:     "Shama",
// 		Email:    "shama@gmail.com",
// 		Password: "shama676",
// 	}

// 	people := []People{Payal, Shama}

// 	for _, person := range people {
// 		go fmt.Println(person.GetName())
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// 	wg.Done()
// 	// Payal.GetName()
// 	// Shama.GetName()

// }
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type People interface {
	GetName() string
	ChangePassword(NewPass string)
}

type Student struct {
	Name     string
	Email    string
	Password string
	Rollno   int
}

func (S Student) GetName() string {
	return fmt.Sprintf("Student Name is : %s", S.Name)
}

func (S Student) ChangePassword(NewPass string) {
	S.Password = NewPass
	fmt.Println("New Password is : ", S.Password)
}

type Teacher struct {
	Name     string
	Email    string
	Password string
}

func (T Teacher) GetName() string {
	return fmt.Sprintf("Teacher Name is : %s", T.Name)
}

func (T Teacher) ChangePassword(NewPass string) {
	T.Password = NewPass
	fmt.Println("Changed password is : ", T.Password)
}

func printName(person People) {
	defer wg.Done()
	fmt.Println(person.GetName())
}

func main() {
	// Created one instance
	Payal := Student{
		Name:     "Payal Sharma",
		Email:    "payalsharma@gmail.com",
		Password: "6778bsu",
		Rollno:   112327,
	}

	Shama := Teacher{
		Name:     "Shama",
		Email:    "shama@gmail.com",
		Password: "shama676",
	}

	people := []People{Payal, Shama}

	for _, person := range people {
		wg.Add(1)
		go printName(person)
	}
	wg.Wait()
}
