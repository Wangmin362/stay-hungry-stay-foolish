package _2_binary_search

// https://www.nowcoder.com/practice/d3df40bd23594118b57554129cadf47b

// 搜索区间为[left, right]
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
