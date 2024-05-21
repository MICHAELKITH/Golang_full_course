package main

import "fmt"

func main() {
	option := 10

	switch option {

	case 5:
		fmt.Println("Option 5")
	case 6:
		fmt.Println("Option 6")
	case 7:
		fmt.Println("Option 7")
	case 8:
		fmt.Println("Option 8")
	case 10:
		fmt.Println("Option 10")

	default:{
		fmt.Println("Others ...")
	}
	}
}