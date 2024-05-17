package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.9
	b := 3.6

	c := math.Pow(a , 2)

	fmt.Printf(" The value of c is %.3f \n",  c)
	d := math.Pow(b , 2)

	fmt.Printf(" The value of d is %.3f \n",  d)

	e := math.Sin(a)

	fmt.Printf(" The value of e is %.3f \n",  e)
	f := math.Cos(b)

	fmt.Printf(" The value of f is %.3f \n",  f)

}