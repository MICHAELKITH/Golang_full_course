package main

import "fmt"
var (
	name  string
	email string
	age   int
)
func main() {
	// var myvar  data_type
	//first way to declare a variable
	var str string
	var n, m int
	var mn float64

	// second way to declare
	

	name ="MIKE"
	email = "mike@gmail.com"
	age = 26

	str = "hello world"
	n = 15
	m = 20
	mn = 15.5

	fmt.Println( "value of str", str)
	fmt.Println( "value of n", n)
	fmt.Println( "value of m", m )
	fmt.Println( "value of mn", mn)
	fmt.Println( "value of name", name)
	fmt.Println( "value of email", email)
	fmt.Println( "value of age", age)

	// :=
	//third way to declare a variable
	country := "Kenya "
	phone := 2223523632

	fmt.Println( "value of country", country)
	fmt.Println( "value of phone", phone)
}
