package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    original := x
    remainder := 0
    reversed := 0
    
    for x != 0 {
        remainder = x % 10
        reversed = reversed * 10 + remainder
        x /= 10
    }
    
    return original == reversed
}

<<<<<<< HEAD
func main() {
	numbers := []int{-100, 100, 121, 1221, 456, 3, 0, }	
	for _, x := range numbers {
		result := isPalindrome(x)
		fmt.Printf("%d is palindrome: %v\n", x, result)
	}
=======

func main() {
	numbers := []int{-121, 121, 2345, 3, 0}

	for _, x := range numbers {
		result := isPalindrome(x)
		fmt.Printf("%d is palindromic: %v\n", x, result)
	}

>>>>>>> 7711cdae1d98e1dae13fffdbd4f45be80c573270
}