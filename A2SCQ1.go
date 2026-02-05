package main

import "fmt"

// function using call by value
func changeValue(x int) {
	x = 50
	fmt.Println("Value inside function =", x)
}

func main() {
	num := 10

	fmt.Println("Value before function call =", num)
	changeValue(num)
	fmt.Println("Value after function call =", num)
}
