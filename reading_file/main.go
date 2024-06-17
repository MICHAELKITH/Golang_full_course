package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Writing content:")
	err := WriteFile("You created your first go file")

	if err != nil {
		log.Fatalf("Failed to write our file: %v", err)
	}

	fmt.Println("Reading file")
	content, err := readFile()

	if err != nil {
		log.Fatalf("Failed to write our file: %v", err)
	}
	fmt.Println("File content:")
	fmt.Println(content)
}

func WriteFile(message string) error{

	file, err := os.Create("test.txt")

	if err != nil{
		return err
	}

	defer file.Close()

	_, err = file.WriteString(message)

	if err != nil{
		return err
	}

	fmt.Println("File created successfully")

	return nil
}


func readFile() (string, error){
	bytes, err := os.ReadFile("test.txt")

	if err != nil {
		return "", err
	}

	return string(bytes), err
}