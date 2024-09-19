package _1_array

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/longest-common-prefix/description/?envType=study-plan-v2&envId=top-interview-150

// 直接从最短的字符串开始遍历
func longestCommonPrefix01(strs []string) string {
	minLen := math.MaxInt32
	minLenStr := ""
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			minLenStr = strs[i]
		}
	}

	for minLen > 0 {
		valid := true
		for i := 0; i < len(strs); i++ {
			if strs[i][0:minLen] != minLenStr[0:minLen] {
				valid = false
				break
			}
		}
		if valid {
			return minLenStr[0:minLen]
		}
		minLen--
	}

	return ""
}

func TestLongestCommonPrefix01(t *testing.T) {
	var testdata = []struct {
		strs []string
		want string
	}{
		{strs: []string{"flower", "flow", "flight"}, want: "fl"},
	}

	for _, tt := range testdata {
		get := longestCommonPrefix01(tt.strs)
		if get != tt.want {
			t.Fatalf("strs:%v, want:%v, get:%v", tt.strs, tt.want, get)
		}
	}
}
