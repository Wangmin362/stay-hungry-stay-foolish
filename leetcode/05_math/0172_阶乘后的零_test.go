package _0_basic

import (
	"fmt"
	"testing"
)

// 题目分析：每个尾随零是由因子 10 产生的，而因子 10 由因子 2 和因子 5 组成。由于阶乘中因子 2 的数量通常多于因子 5，
// 因此我们只需要计算因子 5 的数量。

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		// 在 for 循环中，n /= 5 用于找出有多少个 5 作为因子。每次迭代，我们都把 n 除以 5，这样可以找到 5、10、15 等数中的 5 作为因子。
		n /= 5
		count += n
	}
	return count
}

func TestTrailingZeroes(t *testing.T) {
	n := 25 // 你可以根据需要修改这个值
	result := trailingZeroes(n)
	fmt.Println("尾随零的数量:", result)
}
