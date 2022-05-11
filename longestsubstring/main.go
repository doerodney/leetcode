package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
    maxLen := 0

	for i := 0; i < len(s); i++ {
		// Make an empty set:
		set := make(map[byte]int)

		// Iterate bytes from new starting index:
		for j:= i; j < len(s); j++ {
			b := s[j]
			// Try to get j'th byte from set:
			_, ok := set[b]
			if ok {
				// b is already in the set.  Calculate len of set.
				break
			} else {
				// Add byte to set as key:
				set[b]++
				setLen := len(set)
				if setLen > maxLen {
					maxLen = setLen
				}
			}
		}
	}

	return maxLen
}

func main() {
	targets := []string {"hula hoop", "abcabcbb", "bbbbb", "pwwkew" }

	for idx, x := range targets[0] {
		fmt.Printf("Index: %d, byte: %d\n", idx, x)
	}


	for _, s := range targets {
		l := lengthOfLongestSubstring(s)
		fmt.Printf("Length of longest substring in %s is %d\n", s, l)
	}
}