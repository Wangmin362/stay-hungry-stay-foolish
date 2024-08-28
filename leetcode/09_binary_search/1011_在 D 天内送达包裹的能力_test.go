package _9_binary_search

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/description/

// 题目分析：由于船需要运送货物，因此每个货物都必须要保证可以方，因此最小值为max(weights), 最大值为sum(weights)，
// 这样就可以保证在一天就能运走所有的货物，所以传值的载重能力在[max(weights), sum(weights)]之间。

func shipWithinDays(weights []int, days int) int {
	sum := func() int {
		var res int
		for _, w := range weights {
			res += w
		}
		return res
	}
	canShip := func(ship int) bool {
		var cnt int

		var w int // 当前传值装在的货物重量
		idx := 0
		for idx < len(weights) {
			if w+weights[idx] > ship { // 船只的重量加上下一个物品的重量不能超过传值的载重能力，若超过，只能重新运输
				cnt++
				w = weights[idx]
			} else {
				w += weights[idx]
			}
			idx++
		}
		if w > 0 {
			cnt++
		}

		return cnt <= days
	}

	left, right := slices.Max(weights), sum()
	for left <= right {
		mid := left + (right-left)>>1
		can := canShip(mid)
		if can {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func TestShipWithinDays(t *testing.T) {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
