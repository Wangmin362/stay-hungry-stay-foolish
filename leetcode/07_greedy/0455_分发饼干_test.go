package _0_basic

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/assign-cookies/description/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi := len(g) - 1
	var res int
	for si := len(s) - 1; si >= 0 && gi >= 0; si-- {
		for gi >= 0 && s[si] < g[gi] { // 当前饼干满足不了gi小孩
			gi--
		}
		if gi < 0 {
			return res
		}
		res++
		gi--
	}
	return res
}

// 大饼干尽量给胃口最大的孩子
func findContentChildren02(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	res, si := 0, len(s)-1
	for gi := len(g) - 1; gi >= 0; gi-- {
		if s[si] >= g[gi] { // 说明当前的饼干可以满足这个孩子
			si--
			res++
			if si < 0 {
				break
			}
		}
	}

	return res
}

func TestFindContentChildren(t *testing.T) {
	var testdata = []struct {
		g    []int
		s    []int
		want int
	}{
		{g: []int{1, 2, 3}, s: []int{3}, want: 1},
	}
	for _, tt := range testdata {
		get := findContentChildren02(tt.g, tt.s)
		if tt.want != get {
			t.Fatalf("g:%v, s:%v, want:%v, get:%v", tt.g, tt.s, tt.want, get)
		}
	}
}
