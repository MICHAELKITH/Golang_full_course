package main

import (
	"fmt"
	"time"
)

type Employee struct {

	//field

	id      int
	name    string
	country string
	created time.Time
}

func main() {
	//declare variable

	var emp Employee
	newEmp := new(Employee)

	//set values

	emp.id = 1
	emp.name = "name"
	emp.country = "Kenya"
	emp.created = time.Now()

	newEmp.id = 2
	newEmp.name = "David"
	newEmp.country = "Nigeria"
	newEmp.created = time.Now()
	
	fmt.Println(emp.id)
	fmt.Println(emp.name)
	fmt.Println(emp.country)
	fmt.Println(emp.created)
	fmt.Println("=====")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
}