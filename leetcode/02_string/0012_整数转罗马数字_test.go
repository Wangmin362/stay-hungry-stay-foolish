package _1_array

import (
	"fmt"
	"strings"
	"testing"
)

func intToRoman(num int) string {
	type peer struct {
		roman string
		val   int
	}
	m := []peer{
		{val: 1000, roman: "M"},
		{val: 900, roman: "CM"},
		{val: 500, roman: "D"},
		{val: 400, roman: "CD"},
		{val: 100, roman: "C"},
		{val: 90, roman: "XC"},
		{val: 50, roman: "L"},
		{val: 40, roman: "XL"},
		{val: 10, roman: "X"},
		{val: 9, roman: "IX"},
		{val: 5, roman: "V"},
		{val: 4, roman: "IV"},
		{val: 1, roman: "I"},
	}

	var res strings.Builder
	idx := 0
	for num > 0 {
		for m[idx].val > num {
			idx++
		}
		cnt := num / m[idx].val
		num %= m[idx].val
		for i := 0; i < cnt; i++ {
			res.WriteString(m[idx].roman)
		}
	}

	return res.String()
}

func TestIntToRoman(t *testing.T) {
	fmt.Println(intToRoman(3456))
}
