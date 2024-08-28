package _9_binary_search

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/minimum-time-to-repair-cars/description/

// 题目分析：要把车修好，最少是min(ranks)分钟，当然，这个值肯定不行，最多需要max(ranks)*cars^2，也就是说把所有的
// 都交给能力值最大的修理工修理，那么它锁消耗的时间是最大的。那么时间的取值为[min(ranks), max(ranks)*cars^2]
// 只需要保证在时间t之下所有工人可以修理的数量大于等于cars就可以了

func repairCars(ranks []int, cars int) int64 {
	sum := func(t int) int { // 给定的时间t，所有工人一共可以修理的车辆
		var res int
		for _, r := range ranks {
			res += int(math.Sqrt(float64(t) / float64(r)))
		}
		return res
	}

	left, right := slices.Min(ranks), slices.Max(ranks)*cars*cars
	for left <= right {
		mid := left + (right-left)>>1
		totalCars := sum(mid)
		if totalCars >= cars {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func TestRepairCars(t *testing.T) {
	fmt.Println(repairCars([]int{1, 2, 2, 1, 2, 3, 2, 2, 2}, 4))
}
