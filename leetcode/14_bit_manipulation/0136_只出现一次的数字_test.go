package _0_basic

// https://leetcode.cn/problems/single-number/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}

	return res
}
