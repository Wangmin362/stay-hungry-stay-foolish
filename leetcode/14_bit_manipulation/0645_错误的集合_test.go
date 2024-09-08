package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/convert-a-number-to-hexadecimal/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func findErrorNums(nums []int) []int {
	c := make(map[int]int, len(nums))
	for _, num := range nums {
		c[num]++
	}
	n1, n2 := 0, 0
	for k, v := range c {
		if v == 2 {
			n1 = k
			break
		}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := c[i]; !ok {
			n2 = i
			break
		}
	}
	return []int{n1, n2}
}

func TestFindErrorNums(t *testing.T) {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
