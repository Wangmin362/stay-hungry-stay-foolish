package _0_basic

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] { // 合并区间
			intervals[i][0] = intervals[i-1][0]                       // 使用前一个区间的左端点，因为前一个区间的左端点一定小于等于当前区间的左端点
			intervals[i][1] = max(intervals[i][1], intervals[i-1][1]) // 右端点取两个区间的最大值
		} else { // 没有重叠
			res = append(res, intervals[i-1])
		}
	}
	res = append(res, intervals[len(intervals)-1])
	return res
}
