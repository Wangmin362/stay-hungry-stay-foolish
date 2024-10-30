package _0_basic

import (
	"fmt"
	"testing"
)

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}
		stack = append(stack, i)
	}
	return res
}

func TestDailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
