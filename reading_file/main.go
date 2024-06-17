package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Write data into a file
	fmt.Println("Writing data into a file")
	err := writeFile("Welcome to Go!")
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	// Read data from the file
	fmt.Println("Reading data from the file")
	content, err := readFile()
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println("File content:")
	fmt.Println(content)
}

func writeFile(message string) error {
	file, err := os.Create("testgo.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	fmt.Println("File created successfully")
	return nil
}

func readFile() (string, error) {
	bytes, err := os.ReadFile("testgo.txt")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
