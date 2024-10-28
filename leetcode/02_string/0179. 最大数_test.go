package _1_array

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		s1, s2 := strs[i], strs[j]
		idx := 0
		for ; idx < len(s1) && idx < len(s2); idx++ {
			if s1[idx] > s2[idx] {
				return true
			} else if s1[idx] < s2[idx] {
				return false
			}
		}
		if len(s1) == len(s2) {
			return true
		}

		a1 := s1 + s2
		a2 := s2 + s1
		return a1 > a2
	})

	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(strs[i])
	}
	ans := res.String()
	idx := 0
	for idx < len(ans) && ans[idx] == '0' {
		idx++
	}
	if idx >= len(ans) {
		return "0"
	}

	return ans[idx:]
}
