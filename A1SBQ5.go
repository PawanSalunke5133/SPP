package main

import "fmt"

func main() {
	p := new(int)

	*p = 50

	fmt.Println("Value stored:", *p)
	fmt.Println("Address:", p)
}
