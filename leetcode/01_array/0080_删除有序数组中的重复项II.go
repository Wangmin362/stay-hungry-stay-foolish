package _1_array

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow, fast := 0, 0
	cnt := 0
	for fast < len(nums) {
		if fast == 0 || (fast > 0 && nums[fast] == nums[slow-1]) {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
			fast++
		} else {
			fast++
		}
	}
	return slow
}

func TestRemoveDuplicates(t *testing.T) {
	root := removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 3})
	fmt.Println(root)
}
