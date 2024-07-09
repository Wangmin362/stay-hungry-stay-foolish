package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/roman-to-integer/description/
func romanToInt(s string) int {
	imap := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
		"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900,
	}

	sum := 0
	start, end := 0, len(s)-1
	for start <= end {
		if start+1 > end {
			fmt.Println(s[start : start+1])
			sum += imap[s[start:start+1]]
			start += 1
		} else if val, ok := imap[s[start:start+2]]; ok {
			fmt.Println(s[start : start+2])
			sum += val
			start += 2
		} else {
			fmt.Println(s[start : start+1])
			sum += imap[s[start:start+1]]
			start += 1
		}
	}

	return sum
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("MCMXCIV"))
}
