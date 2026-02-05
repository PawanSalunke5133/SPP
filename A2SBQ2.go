package main

import "fmt"

// function with named return variables
func calculate(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	s, p := calculate(x, y)

	fmt.Println("Sum =", s)
	fmt.Println("Product =", p)
}
