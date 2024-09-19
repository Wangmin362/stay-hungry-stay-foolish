package _1_array

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/integer-to-roman/description/?envType=study-plan-v2&envId=top-interview-150

func intToRoman(num int) string {
	m := map[int]string{
		1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM",
	}
	candidate := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := strings.Builder{}

	idx := 0
	for num > 0 {
		if num < candidate[idx] { // 找到一个可以整数的数字
			idx++
			continue
		}
		consult := num / candidate[idx]
		num %= candidate[idx]
		for i := 0; i < consult; i++ {
			res.WriteString(m[candidate[idx]])
		}
	}

	return res.String()
}

func TestIntToRoman(t *testing.T) {
	var testdata = []struct {
		num  int
		want string
	}{
		{num: 3749, want: "MMMDCCXLIX"},
		{num: 58, want: "LVIII"},
		{num: 1994, want: "MCMXCIV"},
	}

	for _, tt := range testdata {
		get := intToRoman(tt.num)
		if get != tt.want {
			t.Fatalf("num:%v, want:%v, get:%v", tt.num, tt.want, get)
		}
	}
}
