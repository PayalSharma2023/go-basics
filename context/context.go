package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("context in golang")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	go PerformTask(ctx)

	select {
	case <- ctx.Done():
		fmt.Println("Task cancelled")
	}
}

func PerformTask(ctx context.Context){
	select {
	case <- time.After(4*time.Millisecond):
		fmt.Println("Performing Task")
	}
}