package _1_array

import (
	"fmt"
	"strconv"
	"testing"
)

// 解题思路：使用双栈来实现，一个栈保存数字，一个栈保存操作数，括号也看成是一个操作数
func calculate(s string) int {
	ops, nums := make([]byte, 0, len(s)), make([]int, 0, len(s))
	nums = append(nums, 0) // 防止第一个数字出现负数

	idx := 0
	for idx < len(s) {
		switch {
		case s[idx] == '+' || s[idx] == '-' || s[idx] == '(':
			ops = append(ops, s[idx])
		case s[idx] >= '0' && s[idx] <= '9':
			start := idx
			for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
				idx++
			}
			n, _ := strconv.Atoi(s[start:idx])
			nums = append(nums, n)
			idx--
		case s[idx] == ')': // 开始计算数字
			for ops[len(ops)-1] != '(' {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				n1, n2 := nums[len(nums)-2], nums[len(nums)-1]
				nums = nums[:len(nums)-2]
				switch op {
				case '+':
					nums = append(nums, n1+n2)
				case '-':
					nums = append(nums, n1-n2)
				}
			}
			ops = ops[:len(ops)-1] // 弹出右括号
		default:
			// skip black
		}
		idx++
	}
	for len(ops) > 0 {
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		n1, n2 := nums[len(nums)-2], nums[len(nums)-1]
		nums = nums[:len(nums)-2]
		switch op {
		case '+':
			nums = append(nums, n1+n2)
		case '-':
			nums = append(nums, n1-n2)
		}
	}

	return nums[len(nums)-1]
}

func TestCalculate(t *testing.T) {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
