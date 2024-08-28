package _9_binary_search

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/submissions/558950056/

// 很自然的，这道题目可以转化为第一次大于target模型，其实就是求第一次大于等于target+1的索引位置
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}

	leftBoarder := func(letters []byte, target byte) byte {
		left, right := 0, len(letters)-1
		for left <= right {
			mid := left + (right-left)>>1
			if letters[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return letters[left]
	}

	return leftBoarder(letters, target+1)
}
