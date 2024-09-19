package _0_basic

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/word-pattern/description/?envType=study-plan-v2&envId=top-interview-150

func wordPattern(pattern string, s string) bool {
	sp := strings.Split(s, " ")
	if len(pattern) != len(sp) {
		return false
	}

	hashmap := make(map[byte]string, len(pattern))
	hashset := make(map[string]struct{})
	for idx := range pattern {
		word, ok := hashmap[pattern[idx]]
		if !ok {
			if _, ok = hashset[sp[idx]]; ok {
				return false
			}

			hashmap[pattern[idx]] = sp[idx]
			hashset[sp[idx]] = struct{}{}
			continue
		}

		if word != sp[idx] {
			return false
		}
	}

	return true
}
func TestWordPattern(t *testing.T) {
	var testdata = []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abba", t: "dog cat cat dog", want: true},
		{s: "abba", t: "dog cat cat fish", want: false},
		{s: "abba", t: "dog dog dog dog", want: false},
	}
	for _, tt := range testdata {
		get := wordPattern(tt.s, tt.t)
		if get != tt.want {
			t.Fatalf("pattern:%v, t:%v, want:%v, get:%v", tt.s, tt.t, tt.want, get)
		}
	}
}
