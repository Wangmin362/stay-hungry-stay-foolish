package _9_binary_search

// 用于搜索第一次大于等于target的索引，其实就是target需要插入的位置
func leftSearchMeta(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1 // [mid+1, right]
		} else {
			right = mid - 1 // [left, mid-1]
		}
	}
	return left
}

// >= 搜索第一次大于等于target的索引： leftSearchMeta(nums, target)
// >  搜索第一次大于target的索引： leftSearchMeta(nums, target+1)，其实就是搜索第一次大于等于target+1的索引位置
// <= 搜索第一次小于等于target的索引：leftSearchMeta(nums, target+1) -1
// <  所有第一次小于target的索引: leftSearchMeta(nums, target) -1
