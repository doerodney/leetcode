package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
    maxLen := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		// Make an empty set:
		set := make(map[rune]int)

		// Iterate runes from new starting index:
		for j:= i; j < len(runes); j++ {
			r := runes[j]
			// Try to get j'th rune from set:
			_, ok := set[r]
			if ok {
				// Found it.  Calculate len of set.
				setLen := len(set)
				if setLen > maxLen {
					maxLen = setLen
				}
				break
			} else {
				// Add rune to set as key:
				set[r]++
			}
		}
		// fmt.Printf("Index: %d, Letter: %s\n", i, string(letter))
	}

	return maxLen
}

func main() {
	targets := []string {"abcabcbb", "bbbbb", "pwwkew" }
	for _, s := range targets {
		l := lengthOfLongestSubstring(s)
		fmt.Printf("Length of longest substring in %s is %d\n", s, l)
	}
}