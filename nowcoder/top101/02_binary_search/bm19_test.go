package _2_binary_search

// https://www.nowcoder.com/practice/fcf87540c4f347bcb4cf720b5b350c76

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
