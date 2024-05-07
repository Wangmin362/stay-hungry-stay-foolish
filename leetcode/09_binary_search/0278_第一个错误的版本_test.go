package _9_binary_search

func isBadVersion(version int) bool {
	return false
}

// 而直接使用二分查找算法，查找区间为[left, right]
func firstBadVersion(n int) int {
	if n <= 0 {
		return -1
	}

	left := 1
	right := n
	target := -1        // 第一个坏版本
	for left <= right { // 当left = right时, [left, right]区间依然时有效的
		middle := left + (right-left)>>1
		if isBadVersion(middle) {
			target = middle
			right = middle - 1 // 由于middle已经是坏版本了，所以第一个版本只可能在中位数左边
		} else {
			left = middle + 1 // middle依然不是换版本，那么坏版本只可能在右边
		}
	}

	return target
}
