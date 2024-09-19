package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) <= 0 {
		return true
	}

	cache := make(map[byte]byte, len(s))
	rcache := make(map[byte]struct{}, len(s))
	for idx := range s {
		// s -> t
		scmapping, ok := cache[s[idx]]
		if ok {
			if scmapping != t[idx] {
				return false
			}
		} else { // 说明目前还没有映射
			// t -> s
			if _, ok = rcache[t[idx]]; ok { // 说明当前字符已经映射
				return false
			}

			cache[s[idx]] = t[idx]
			rcache[t[idx]] = struct{}{}
		}
	}

	return true
}

func TestIsIsomorphic(t *testing.T) {
	var testdata = []struct {
		s    string
		t    string
		want bool
	}{
		{s: "egg", t: "add", want: true},
		{s: "paper", t: "title", want: true},
		{s: "bbbaaaba", t: "aaabbbba", want: false},
		{s: "badc", t: "baba", want: false},
	}
	for _, tt := range testdata {
		get := isIsomorphic(tt.s, tt.t)
		if get != tt.want {
			t.Fatalf("s:%v, t:%v, want:%v, get:%v", tt.s, tt.t, tt.want, get)
		}
	}
}
