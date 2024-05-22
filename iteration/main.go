package main

import (
	"fmt"
	// "slices"
)

func main() {
	//iteration

	//for initial; condtion; inc/dec{}

	// var i int

	// for i = 0; i <= 10; i++ {

	// 	fmt.Println(i)

		//0
		//1
		//2
		//3.....

	// }


	// numbers := []int {10, 1, 4, 7, 9, 20}

	// slices.Sort(numbers)
	// for _, value := range numbers{
	// 	fmt.Println(value)
	// }

	//


	//iteration with while loop 


	// i := 0

	// for i< 10{
	// 	fmt.Println(i)
	// 	i++
	// }

	//break and continue

	i := 0
	j := 5
	fmt.Print("Break")
	for i = 0; i<10; i++{
		if i == 3{
			break
		}

		fmt.Println(i)
	}

fmt.Print("Continue")
	for j = 5; j<10; j++{
		if j == 5{
			continue
		}

		fmt.Println(j)
	}

}