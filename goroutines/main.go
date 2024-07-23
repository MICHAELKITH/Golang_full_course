package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	//Add method

	wg.Add(3)
	fmt.Println("Starting")
	go count("Runner 1",  &wg)
	go count("Runner 2",  &wg)
	go count("Runner 3",  &wg)

	//sleep our program
	wg.Wait()

	fmt.Print("Stopped Running")
}

func count(value string, wg *sync.WaitGroup){
	fmt.Println(value , "STARTED")

	defer wg.Done()

	for i:=1; i<=5; i++{
		fmt.Println(value, i)
	}
}