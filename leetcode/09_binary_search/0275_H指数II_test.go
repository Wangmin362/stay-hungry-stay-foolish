package _9_binary_search

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/h-index-ii/description/

// 题目分析：假设计算的指数为h，那么至少需要有h篇文章被引用了h次数，h的范围为[0, len(citation)]
// 计算citations大于等于h的索引，从而获取数量，如果数量大于等于h,就说明这个h是合法的

func hIndex(citations []int) int {
	leftBoarder := func(h int) int { // 返回指数为h时，大于等于h的文章数量
		left, right := 0, len(citations)-1
		for left <= right {
			mid := left + (right-left)>>1
			if citations[mid] < h {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return len(citations) - left
	}

	left, right := 0, len(citations)
	for left <= right {
		mid := left + (right-left)>>1
		validPeaper := leftBoarder(mid)
		if validPeaper >= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func TestHIndex(t *testing.T) {
	fmt.Println(hIndex([]int{1, 2, 100}))
}
