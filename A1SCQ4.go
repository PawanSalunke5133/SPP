package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n >= -9 && n <= 9 {
		fmt.Println("Single digit number")
	} else {
		fmt.Println("Not a single digit number")
	}
}
