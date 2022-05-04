package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    var result []int
    dict := make(map[int]int) // key = number, value = index

	// Create map of values:
	for value, key := range(nums) {
		dict[key] = value
	}

	// Iterate through the list of numbers.  
	// Find the pair whose sum is the target value.
	for idx, number := range(nums) {
		key := target - number

		if value, ok := dict[key]; ok {
			if idx != value {
				result = append(result, idx, value)
				break
			}
		}
	}
    return result
}


func main() {
	nums := []int {3,2,4}
	target := 6

	result := twoSum(nums, target)
	fmt.Printf("Result: %v\n", result)
}
