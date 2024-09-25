package _0_basic

import "testing"

// https://leetcode.cn/problems/interleaving-string/description/?envType=study-plan-v2&envId=top-interview-150

func isInterleave(s1 string, s2 string, s3 string) bool {
	return false
}

func TestIsInterleave(t *testing.T) {
	var testdata = []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{s1: "", s2: "", s3: "", want: true},
	}

	for _, tt := range testdata {
		get := isInterleave(tt.s1, tt.s2, tt.s3)
		if get != tt.want {
			t.Fatalf("s1:%v, s2:%v, s3:%v, want:%v, get:%v", tt.s1, tt.s2, tt.s3, tt.want, get)
		}
	}
}
