package _0_basic

import "sort"

// 思路：其实就是区间问题，看看重合的区间有几个。 先按照区间的左边界从小到大排序。然后更新区间的有边界
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var res int
	for i := 0; i < len(points); i++ {
		if i == 0 { // 肯定需要一只箭
			res++
			continue
		}

		if points[i][0] > points[i-1][1] { // 如果当前区间的左边界比上一个区间的有边界大，说明不重叠
			res++
		} else { // 说明重叠，更新当前区间的有边界为两个区间有边界的最小值
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return res
}
