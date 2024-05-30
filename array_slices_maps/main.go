package main

import "fmt"

func main(){
//array
	//short syntax
	// x := [6]float64{
	// 	98,56,50,90,83,76,
	// }
	


	//iteration
	// var total float64 = 0

	// for _, value := range x {
	// 	total += value
	// }
		
	

	// fmt.Println( total/float64(len(x)) )

	

	//slices
	// var x [] float64

	//shortform

	// x:= make([]float64, 5, 10)

	//Low and high expression

	// x:= arr[0:5]



	// slice := []int{1,2,3}

	// slice1 := append(slice, 4,5,6)

	// fmt.Println(slice, slice1)

	//copy

	// slice1 := []int{1,2,3}

	// slice2 := make([]int, 2)

	// copy(slice2, slice1)

	// fmt.Println(slice1, slice2)


	//Maps

	elements := make(map[string]string)
	elements["H"]="House"
	elements["C"]="Cat"
	elements["D"]="Dog"
	elements["M"]="Mouse"
	elements["U"]="Dog"
	elements["B"]="Mouse"


	if name, ok := elements["K"]; ok{
		fmt.Println(name, ok)
	}
}
