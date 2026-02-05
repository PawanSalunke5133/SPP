package main

import "fmt"

// function returning multiple values
func operations(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	s, d := operations(x, y)

	fmt.Println("Sum =", s)
	fmt.Println("Difference =", d)
}
