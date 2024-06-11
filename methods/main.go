package main

import (
	"fmt"
	"math"
)

//struct
type Circle struct{

	x,y int
	r float64
}

type Author struct{
	name string
	branch string 
	partcles int
}


func (c Circle) display(){
	fmt.Printf("x=%d, y=%d, r=%.2f", c.x, c.y, c.r)
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.r, 2)
}
func (c Circle) Circumference() float64{
	return 2*math.Pi * c.r
}


func (a *Author) show(abranch string){
	(*a).branch = abranch
}
func main() {
	// shape := Circle{10,5,3.8}
	// shape.display()
	// a:=shape.area()
	// C:= shape.Circumference()
	

	// fmt.Printf("Area= %.2f\n", a)
	// fmt.Printf("Circumference= %.2f\n", C)
	

	//pointers methods 
	res := Author{
		name: "Michael",
		branch: "Nairobi",
	}

	fmt.Println("Author's name \n", res.name)
	fmt.Println("Branch name(Before) \n", res.branch)
	

	p:= &res

	p.show("New York")

	fmt.Println("Author's name \n", res.name)
	fmt.Println("Branch name(After) \n", res.branch)

}






