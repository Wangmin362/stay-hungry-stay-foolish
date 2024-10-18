package _1_array

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/longest-common-prefix/description/?envType=study-plan-v2&envId=top-interview-150

func longestCommonPrefix(strs []string) string {
	// 找到最短的字符串
	minLen, minLenStr := math.MaxInt32, ""
	for i := 0; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
			minLenStr = strs[i]
		}
	}
	if minLenStr == "" { // 有空字符串，肯定就是空字符串
		return minLenStr
	}

	for i := minLen - 1; i >= 0; i-- {
		ptn := minLenStr[:i+1]
		match := true
		for j := 0; j < len(strs); j++ {
			if strs[j][:i+1] != ptn {
				match = false
				break
			}
		}
		if match {
			return ptn
		}
	}

	return ""
}

// 灵神解题思路：竖着向下看，如果所有字符相同，就往后移动
func longestCommonPrefix02(strs []string) string {
	for idx := range strs[0] {
		ch := strs[0][idx]
		for i := 0; i < len(strs); i++ {
			if idx >= len(strs[i]) || strs[i][idx] != ch {
				return strs[0][:idx]
			}
		}
	}
	return strs[0]
}

func TestLongestCommonPrefix01(t *testing.T) {
	var testdata = []struct {
		strs []string
		want string
	}{
		{strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{strs: []string{"dog", "racecar", "car"}, want: ""},
		{strs: []string{"d"}, want: "d"},
	}

	for _, tt := range testdata {
		get := longestCommonPrefix02(tt.strs)
		if get != tt.want {
			t.Fatalf("strs:%v, want:%v, get:%v", tt.strs, tt.want, get)
		}
	}
}
