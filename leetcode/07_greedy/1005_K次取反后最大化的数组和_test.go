package _0_basic

import (
	"fmt"
	"sort"
	"testing"
)

// 贪心思路：先排序，如果有负数，就先把最小的负数反转，反转完成之后，找到最小值，反转剩余的次数
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var i int
	for i = 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 { // 有负数先把负数反转
			nums[i] = -nums[i]
			k--
		} else {
			break
		}
	}

	if k%2 == 1 { // 说明还需要反转，找到最小的数，反转剩余次数
		minIdx := i
		if minIdx == len(nums) {
			minIdx -= 1
		} else if i > 0 && nums[i] > nums[i-1] {
			minIdx = i - 1
		}
		nums[minIdx] = -nums[minIdx]
	}

	var res int
	for i = 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}
func TestAbc(t *testing.T) {
	//fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	fmt.Println(largestSumAfterKNegations([]int{-4, -2, -3}, 4))
}
