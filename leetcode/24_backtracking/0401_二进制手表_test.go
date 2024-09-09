package _1_array

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/binary-watch/description/?envType=problem-list-v2&envId=backtracking&difficulty=EASY

// 其实就是0-9选取k个数字
func readBinaryWatch(k int) []string {
	var backtracking func(start int)

	var res []string
	var path []int
	backtracking = func(start int) {
		if len(path) == k {
			hour := 0
			minute := 0
			for _, p := range path {
				if p <= 5 {
					minute += 1 << p
					if minute > 59 {
						return
					}
				} else {
					hour += 1 << (p - 6)
					if hour > 12 {
						return
					}
				}
			}
			res = append(res, fmt.Sprintf("%d:%02d", hour, minute))
			return
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	sort.Strings(res)
	return res
}

func TestReadBinaryWatch(t *testing.T) {
	fmt.Println(readBinaryWatch(2))
}
