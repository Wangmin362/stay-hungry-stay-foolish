package _1_array

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/gray-code/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func grayCode(n int) []int {
	var backtracking func(num int)

	var res []int
	var path []int
	getRest := false
	length := int(math.Pow(float64(2), float64(n)))
	cache := make([]bool, length)
	backtracking = func(num int) {
		if getRest {
			return
		}
		if len(path) == length {
			getRest = true
			tmp := make([]int, length)
			copy(tmp, path)
			res = tmp
			return
		}

		if cache[num] {
			return
		}

		for i := 0; i < n; i++ {
			next := num
			if next&(1<<i) == 0 {
				next |= 1 << i
			} else {
				mask := 1 << i
				mask = ^mask
				next &= mask
			}

			cache[num] = true
			path = append(path, num)
			backtracking(next)
			path = path[:len(path)-1]
			cache[num] = false
		}
	}

	backtracking(0)
	return res
}

func TestGrayCode(t *testing.T) {
	fmt.Println(grayCode(3))
}
