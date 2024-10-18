package _0_basic

import (
	"sort"
	"testing"
)

// 贪心，尽可能的把当前船只装满，优先装体重大的人
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	var res int
	for left <= right {
		capa := limit - people[right] // 每条船一定可以装最后一个人
		right--
		if capa >= people[left] { // 说明当前船只还可以再装一个人
			left++
		}
		res++
	}

	return res
}

func TestName(t *testing.T) {
	t.Log(numRescueBoats([]int{1, 2}, 3))
}
