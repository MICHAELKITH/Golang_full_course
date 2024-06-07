package main

import "fmt"

type Name struct {
	id      int
	name    string
	country string
}

func main() {
	var name Name

	newName := new(Name)

	newName.id = 5
	name.id = 5

	fmt.Println("name id", name.id)
	fmt.Println("name id", newName.id)
}
