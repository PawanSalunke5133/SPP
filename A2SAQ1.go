package main

import "fmt"

// function to add two numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	result := add(x, y)
	fmt.Println("Addition =", result)
}
