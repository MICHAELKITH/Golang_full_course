package main

import "fmt"

// Logic Operators
// && and
// || or ! not

func main(){
var (
	a= 10
	b =15
)

fmt.Println(a>b && a!=b) //false
fmt.Println(a==b && a>b) //false 
fmt.Println(!(a>=b)) //true

}