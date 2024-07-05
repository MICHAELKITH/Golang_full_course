package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello goroutines")

	time.Sleep(time.Millisecond *2)
	fmt.Println("Hello goroutines 2")
}