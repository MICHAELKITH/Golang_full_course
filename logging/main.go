//logging

package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	simpleLogging()
	formattingLogging()
	fileLogging()

}

func simpleLogging(){
	fmt.Println("----------------Simple logging------------")
	log.Println("Hello world")
	log.Println("Simple error occured")
}

func formattingLogging(){
	fmt.Println("----------------Simple logging------------")

	var warning *log.Logger

	warning = log.New(
		os.Stdout,
		"Warning",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	warning.Println("This a warning one")
	warning.Println("This a warning two")
	


}
func fileLogging(){
	fmt.Println("----------------Simple logging------------")

	file, err := os.OpenFile("test.txt",
							os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil{
		fmt.Println("failed to open file")
		return
	}
	var logFile *log.Logger

	logFile = log.New(
		file, 
		"APP:",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	logFile.Println("This error message one")
	logFile.Println("This error message two")
	logFile.Println("This error message three")

	fmt.Println("Done")



}