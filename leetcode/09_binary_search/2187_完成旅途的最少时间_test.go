package _9_binary_search

import "slices"

// https://leetcode.cn/problems/minimum-time-to-complete-trips/description/

// 题目分析：最小的时间为1，最大的时间为max(time)*totalTrips，也就是说最多我消费max(time)*totalTrips的时间
// 就可以完成这个旅途的次数。那么时间的范围取值为[1, max(time)*totalTrips], 每个公交车在t时间下能够完成的次数为
// t/time[i]，并且是向下取整。那么总共的次数就是所有的和。根据除法单调性原理，选取的时间越小，total越小，选取的t
// 越大，那么total越大，我只需要找到一个时刻t-1, total(t-1)< totalTrips, 并且total(t) >= totalTrips，那么
// 这个t就是合法的

func minimumTime(time []int, totalTrips int) int64 {
	sum := func(t int) int {
		var total int
		for _, busTime := range time {
			total += t / busTime
		}
		return total
	}

	left, right := 1, slices.Max(time)*totalTrips
	for left <= right {
		mid := left + (right-left)>>1
		total := sum(mid)
		if total >= totalTrips {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}
