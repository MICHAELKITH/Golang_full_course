package main

import "fmt"

func main() {

	names := []string{"mike", "joan", "bro"}
	for _, n := range names{
		fmt.Printf("name are %v\n", n)
	}

}