package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/single-number-iii/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=MEDIUM

func singleNumberIII(nums []int) []int {
	c := make(map[int]int)
	for _, num := range nums {
		cnt := c[num]
		if cnt == 0 {
			c[num]++
		} else {
			delete(c, num)
		}
	}
	var res []int
	for k := range c {
		res = append(res, k)
	}

	return res
}

// 参考：灵茶山艾府
func singleNumberIII02(nums []int) []int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	var i int
	for i = 0; i < 32; i++ { // 找到第一个1
		if res&0x1 == 1 {
			break
		}
		res >>= 1
	}

	n := 1 << i
	a1, a2 := []int{}, []int{}
	for _, num := range nums {
		if num&n == 0 {
			a1 = append(a1, num)
		} else {
			a2 = append(a2, num)
		}
	}

	n1 := a1[0]
	for i := 1; i < len(a1); i++ {
		n1 ^= a1[i]
	}

	n2 := a2[0]
	for i := 1; i < len(a2); i++ {
		n2 ^= a2[i]
	}

	return []int{n1, n2}
}

// 参考：灵茶山艾府
func singleNumberIII03(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	lowerbit := xor & -xor
	res := []int{0, 0}
	for _, num := range nums {
		if num&lowerbit == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}

	return res
}

func TestSingleNumberIII(t *testing.T) {
	fmt.Println(singleNumberIII03([]int{1, 2, 1, 3, 2, 5}))
}
