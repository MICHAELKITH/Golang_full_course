package main

import (
	"fmt"
	"math"
)

func main() {
//  foo()
circle_area(2.5)

calculate(2, 5.5, 2.3)
advanced_calculate(3,5,6.6)

compute(1.4,5.5,6.3, "michael")

closure_func("Elon")

}
//simple funct
// func foo() {
// 	fmt.Println("Foo was called")
// }

//functions with parameters

func circle_area(r float64){
	area  := math.Pi * math.Pow(r,2)


	fmt.Printf("Circle area with radius %.2f is %.2f \n",  r, area)
}


func calculate(a, b, c float64){
	result := a*b*c


	fmt.Printf("a=%.2f b=%.2f c=%.2f and result= %.2f\n", a, b , c , result)
}


//return single
func advanced_calculate(a, b, c float64) float64{

	result := a *b*c
	return result

	
}

//mutilple returning
func compute(a, b, c float64 , name string) (float64, float64, string){

	result1 := a *b*c
	result2 := a +b+c
	result3 := result1 +result2

	newName := fmt.Sprintf("%s value = %.2f", name , result3)
	return result1 , result2, newName

	
}


//Unlimited 
func add (numbers ...int) int{
	result := 0

	for _, v := range numbers{
		result += v
	}

	return result
}

//closure

func closure_func(name string){
	hoo := func(a,b int){
		result := a*b

		fmt.Printf("hoo()= %d \n", result)
	}
	joo := func(a,b int) int{
		result := a*b+a
		return result
		
	}

	fmt.Printf("Closure_func %s was called\n", name)

	hoo(2,3)
	val := joo(4,5)

	fmt.Printf("Val from joo() = %d \n", val)
}

//recursion 

//Finobacci



