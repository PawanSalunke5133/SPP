package main

import "fmt"

// function to check palindrome
func isPalindrome(n int) bool {
	original := n
	reverse := 0

	for n > 0 {
		reverse = reverse*10 + n%10
		n = n / 10
	}

	return original == reverse
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println("Number is Palindrome")
	} else {
		fmt.Println("Number is Not Palindrome")
	}
}
