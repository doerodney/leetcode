package main

import (
	"fmt"
)

func romanToInt(s string) int {
	intByRoman := map[string]int {
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
		"IV" : 4,
		"IX" : 9,
		"XL" : 40,
		"XC" : 90,
		"CD" : 400,
		"CM" : 900,
	}

	key := ""
	sum := 0
	value := 0
	ok := false
	i := 0
	for i < len(s) {
		ok = false
		value = 0
		// Lookup the next two characters:
		if (i + 1) < len(s) {
			key = string(s[i:i + 2])
			value, ok = intByRoman[key]
		}
		// If ok, then two character lookup worked.
		if ok {
			i += 2
		} else {
			// Get the value of the next character.
			key = string(s[i])
			value, ok = intByRoman[key]
			i += 1
		}
		sum += value		 
	}

	return sum
}


func main() {
	romans := []string{"I", "II", "III", "IV", "V", "CMXCIV"}
	for _, s := range romans {
		d := romanToInt(s)
		fmt.Printf("%s = %d\n", s, d)
	}
}