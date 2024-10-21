package _9_binary_search

// 参考灵神视频：https://www.bilibili.com/video/BV1QK411d76w/?vd_source=d039ae9ec8b71e411a906e821301b7ac

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-2
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
