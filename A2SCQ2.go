package main

import (
	"fmt"
	"os"
)

func main() {
	// create a file
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	// close file using defer
	defer file.Close()

	// write to file
	file.WriteString("Hello World")

	fmt.Println("File created and data written successfully")
}
