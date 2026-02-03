package main

import "fmt"

func main() {
	var s1, s2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&s1)
	fmt.Print("Enter second string: ")
	fmt.Scan(&s2)

	p1 := &s1
	p2 := &s2

	result := *p1 + *p2

	fmt.Println("Concatenated string:", result)
}
