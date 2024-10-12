package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// 贪心思路：如果当前的累加和是负数，就抛弃，因为负数加上后边的数一定会让累加后变小，只有当之前的累加和为正数的时候才有意义
// 重点是退回这种贪心的思路，如果累加和为负数，后续的累加和一定会更小，所以要抛弃，只有当为正数的时候才需要计算
func maxSubArray(nums []int) int {
	maxSum, currSum := math.MinInt32, 0
	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		maxSum = max(maxSum, currSum)
		if currSum < 0 {
			currSum = 0 // 对其之前的累加和
		}
	}

	return maxSum
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
