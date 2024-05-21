package main

import "fmt"

// operators
// == is equal to ,  =! not equal
// > is greater than
// < is less than
// >= is greater than or equal to
// <= is less than or equal to
func main (){

	var (
		a= 20
		b= 38
	)

	fmt.Println(a>b) //false 
	fmt.Println(a<b) //true
	fmt.Println(a >= b) //false
	fmt.Println(a != b) //true
	fmt.Println(a<=b) //true

}