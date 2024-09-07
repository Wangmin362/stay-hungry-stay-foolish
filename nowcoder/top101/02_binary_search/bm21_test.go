package _2_binary_search

// https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5

func minNumberInRotateArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
