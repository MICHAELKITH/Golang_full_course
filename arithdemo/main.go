package main

import "fmt"

func main() {
	var a, b int

	a = 9
	b = 20

	c := a + b
	//%d means decimal integer
	fmt.Printf( "%d + %d = %d \n",  a, b , c)
	d := b - a
	//%d means decimal integer
	fmt.Printf( "%d - %d = %d \n",  b, a , d)

	e := float32(b) / float32(a)
	//%d means decimal integer
	fmt.Printf( "%d / %d = %.3f \n",  b, a , e)
}