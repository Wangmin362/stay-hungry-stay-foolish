package _1_array

import (
	"fmt"
	"testing"
)

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= k {
			if sum == k {
				res++
			}
			sum -= nums[left]
			left++
		}
	}

	return res
}

func TestSubarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}
