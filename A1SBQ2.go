package main

import "fmt"

func main() {
	var rows int
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		num := 1
		for space := 0; space < rows-i; space++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
