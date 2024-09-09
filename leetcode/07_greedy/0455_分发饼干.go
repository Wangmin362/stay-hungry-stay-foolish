package _0_basic

import "sort"

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
