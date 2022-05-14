package main

import (
	"fmt"
)


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }  
    
    reversed := 0 
    remainder := 0 
    original := x 
    
    for x != 0 {
        remainder = x % 10
        reversed = reversed * 10 + remainder
        x /= 10
    }
    
    return original == reversed
}


func main() {
	numbers := []int{-121, 121, 2345, 3, 0}

	for _, x := range numbers {
		result := isPalindrome(x)
		fmt.Printf("%d is palindromic: %v\n", x, result)
	}

}