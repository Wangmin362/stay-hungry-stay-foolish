package _1_array

import (
	"fmt"
	"strconv"
	"testing"
)

func addStrings(num1 string, num2 string) string {
	var res string
	overflow := byte(0)
	i1, i2 := len(num1)-1, len(num2)-1
	for i1 >= 0 || i2 >= 0 {
		v1, v2 := byte(0), byte(0)
		if i1 >= 0 {
			v1 = num1[i1] - '0'
			i1--
		}
		if i2 >= 0 {
			v2 = num2[i2] - '0'
			i2--
		}
		sum := v1 + v2 + overflow
		if sum > 9 {
			sum %= 10
			overflow = 1
		} else {
			overflow = 0
		}
		res = strconv.Itoa(int(sum)) + res
	}
	if overflow == 1 {
		return "1" + res
	}
	return res
}

func TestAddStrings(t *testing.T) {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("99999", "1"))
}
