package main

import "fmt"

func main() {

	defer func(){
		if r := recover()
		r != nil {
			fmt.Println("System recovered\n", r)
			fmt.Println("This line will  be executed!")
			errHandle()
		}
	}()

	myPanic()
	

}

func errHandle() {

	defer fmt.Println("Error handling 1")
	defer fmt.Println("Error handling 2")
	defer fmt.Println("Error handling 3")
}

func myPanic() {
	panic("problem occured")
}
