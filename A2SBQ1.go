package main

import "fmt"

// function to swap values using pointers
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	swap(&x, &y)

	fmt.Println("After swapping:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
