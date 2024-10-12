package _0_basic

import (
	"fmt"
	"sort"
	"testing"
)

// 思路：显然，肯定需要先把区间按照左边从小到大排序，如果区间有重叠，就需要移除区间较长的那个，这样可以移除最少的区间
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] { // 如果左端点相同，那么需要按照右区间按照从小到大排序
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	if len(intervals) < 2 {
		return 0
	}

	var res int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] { // 说明没有重叠
			// pass
		} else { // 说明有重叠，删除当前区间
			res++
			// TODO 删除区间这里需要注意，我删除不一定是当前区间，也有可能是上一个区间。那么到底删除那个区间呢？
			// 其实我只需要删除有边界区间最长的那个区间，保留有边界区间最小的区间，这样就可以尽可能的保证区间没有重叠，从而最少的删除区间
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}
	return res
}

func TestEraseOverlapIntervals(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
}
