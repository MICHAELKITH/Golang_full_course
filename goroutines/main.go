package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex
	//Add method

	wg.Add(3)
	fmt.Println("Starting")
	go count("Runner 1",  &wg, &m)
	go count("Runner 2",  &wg, &m)
	go count("Runner 3",  &wg, &m)

	//sleep our program
	wg.Wait()

	fmt.Print("Stopped Running")
}

func count(value string, wg *sync.WaitGroup, m *sync.Mutex){
	fmt.Println(value , "STARTED")

	defer wg.Done()

	for i:=1; i<=10; i++{
		m.Lock()
		fmt.Println(value, i)
		m.Unlock()
	}
}