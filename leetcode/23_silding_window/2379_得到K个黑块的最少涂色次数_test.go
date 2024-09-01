package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/

// 题目分析：想要获取连续k个黑色块的数量，其实就是统计窗口为k的字符串中有多少个白色块

func minimumRecolors(blocks string, k int) int {
	ans, sum := k, 0 // 至少需要涂k次，因为可能全是白色
	for idx, in := range blocks {
		if in == 'W' {
			sum++
		}

		if idx < k-1 {
			continue
		}

		ans = min(ans, sum)

		out := blocks[idx-k+1]
		if out == 'W' {
			sum--
		}
	}
	return ans
}
func TestMinimumRecolors(t *testing.T) {
	testdata := []struct {
		blocks string
		k      int
		expect int
	}{
		{blocks: "WBBWWBBWBW", k: 7, expect: 3},
	}

	for _, test := range testdata {
		get := minimumRecolors(test.blocks, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.blocks, test.k, test.expect, get)
		}
	}
}
