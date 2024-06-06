package main

import "fmt"

func main(){

	x:= 1

	fmt.Println(x) //print value
	fmt.Println(&x) // print address

	//declare pointer 

	var num *int = &x

	val:= new(int)

	val = &x

	*val = 2


fmt.Println("===pointer num===")
fmt.Println(*num) // print a value of x
fmt.Println(num) // print address of x
fmt.Println("===pointer val===")
fmt.Println(*val) // print a value of x
fmt.Println(val) // print address of x

//working with pointers
var creature string = "shark"

var pointer *string = &creature


fmt.Println("*pointer", pointer) //address
fmt.Println("creature", creature) //shark

fmt.Println("*pointer", *pointer) //shark
fmt.Println("creature", creature) //shark

*pointer = "snake"

fmt.Println("*pointer", *pointer) //snake
fmt.Println("creature", creature)  //snake 
}