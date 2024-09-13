package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/n-th-tribonacci-number/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=EASY

func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	t0, t1, t2 := 0, 1, 1
	// t0, t1, t2, res
	//     t0, t1, t2, res
	for i := 3; i <= n; i++ {
		res := t0 + t1 + t2
		t0 = t1
		t1 = t2
		t2 = res
	}
	return t2
}

func TestTribonacci(t *testing.T) {
	res := tribonacci(10)
	fmt.Println(res)

}
