package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/insert-into-bits-lcci/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func insertBits(N int, M int, i int, j int) int {
	high := N >> (j + 1)
	low := N & (1<<i - 1)
	res := high<<(j+1) | M<<i | low
	return res
}

func TestInsertBits(t *testing.T) {
	fmt.Println(insertBits(1024, 19, 2, 6))
}
