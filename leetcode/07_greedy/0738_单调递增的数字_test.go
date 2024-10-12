package _0_basic

import (
	"fmt"
	"testing"
)

// 这种算法会超时
func monotoneIncreasingDigits(n int) int {
	isValid := func(num int) bool {
		if num < 10 {
			return true
		}
		last := 10 // 初始化一个大点的数字，单个数字不可能大于10
		for num != 0 {
			mod := num % 10
			if mod > last {
				return false
			}
			last = mod
			num /= 10
		}

		return true
	}

	for !isValid(n) {
		n--
	}

	return n
}

// 解题思路：从末尾开始向前遍历，如果当前数字比前面的数字小，那么就把前面的数字减一，相当于借位，与此同时把当前数字以及后面所有的数字设置为9
// 从而保证尽可能大
func monotoneIncreasingDigits02(n int) int {
	digits := make([]int, 0, 20)
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] { // 如果当前的位比前一位小
			// 1. 当前位置以及之后所有的位置都设置位9
			// 2. 前一位需要减一，相当于借位了
			digits[i+1] -= 1
			for j := 0; j <= i; j++ {
				digits[j] = 9
			}
		}
	}

	var res int
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}

	return res
}

func TestMonotoneIncreasingDigits(t *testing.T) {
	fmt.Println(monotoneIncreasingDigits02(554889396))
}
