// WaitGroups
package main

import (
	"fmt"
	"sync"
)


func main ()  {
	var wg  sync.WaitGroup

	for i:=1; i <=5; i++{
		wg.Add(1)

		go func (){
			defer wg.Done()

			student(i)
		}()
	}
	wg.Wait()
}

func student(id int){

	fmt.Printf("Student %d starting\n",  id)


}