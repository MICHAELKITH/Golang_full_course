package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var str string
	var str1 string

	str = "Welcome to 55 Blocks Channel"
	str1 = "Subscribe to 55 Blocks Channel"

	str4 := fmt.Sprintf("%s and %s", str, str1)

	fmt.Println("My concatenated string is:", str4)

	//string to numeric


	str_val := "5"
	// str_val1 := "10"

	var err error

	var int_val int

	int_val, err = strconv.Atoi(str_val)

	if err == nil{
		println("intValue is :", int_val)
	}else{
		fmt.Println(err)
	}


	//Numeric to string

	num1 := 1

	var str_num string

	str_num = fmt.Sprintf("%d", num1)

	fmt.Println(str_num)


	//string parser


	data := "Kenya; NewYork; China; Nigeria; Uganda"

	fmt.Println(data)

	places := strings.Split(data, ";")

	for _, place := range places{
		fmt.Println(place)
	}

	//string len

	check_ln:= len(str)


	fmt.Println(check_ln)


	//ToUpper/ToLower


	upper:= strings.ToUpper(str)
	lower:= strings.ToLower(str)

	fmt.Println(upper, lower)
}