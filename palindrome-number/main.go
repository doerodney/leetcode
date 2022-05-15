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

func main() {
	numbers := []int{-100, 100, 121, 1221, 456, 3, 0, }	
	for _, x := range numbers {
		result := isPalindrome(x)
		fmt.Printf("%d is palindrome: %v\n", x, result)
	}
}