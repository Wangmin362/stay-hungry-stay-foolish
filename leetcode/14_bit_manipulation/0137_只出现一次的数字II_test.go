package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/single-number-ii/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=MEDIUM

// 统计每个[0,31]个bit出现的次数，然后取模，就能知道这个bit真正的值

func singleNumberII(nums []int) int {
	var res int32
	for i := 0; i <= 31; i++ {
		cnt := int32(0)
		for _, num := range nums {
			cnt += (int32(num) >> i) & 0x1
		}
		res |= (cnt % 3) << i
	}
	return int(res)
}

func TestSingleNumberII(t *testing.T) {
	fmt.Println(singleNumberII([]int{2, 2, 3, 2}))
}
