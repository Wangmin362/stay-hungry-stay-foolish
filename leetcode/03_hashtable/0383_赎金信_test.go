package _0_basic

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/ransom-note/

// 解题思路：
// 解法一：由于限定了字符串为应为小写字母，因此可以使用数组作为map来解决此问题
// 解法二：通用一点的化，还是应该使用map，应为字符串可能为Unicode

func canConstruct(ransomNote string, magazine string) bool {
	engMap := make([]rune, 26)
	for _, mChar := range magazine {
		cnt := engMap[mChar-'a'] // 默认数组初始化为0
		engMap[mChar-'a'] = cnt + 1
	}

	for _, rChar := range ransomNote {
		cnt := engMap[rChar-'a']
		if cnt > 0 {
			engMap[rChar-'a'] = cnt - 1
		} else {
			return false
		}
	}

	return true
}

func canConstruct03(ransomNote string, magazine string) bool {
	cache := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		cache[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		cnt := cache[ransomNote[i]-'a']
		if cnt >= 1 {
			cache[ransomNote[i]-'a']--
		} else {
			return false
		}
	}

	return true
}

func TestCanConstruct(t *testing.T) {
	var testdata = []struct {
		s      string
		t      string
		expect bool
	}{
		{s: "a", t: "b", expect: false},
		{s: "aa", t: "ab", expect: false},
		{s: "aa", t: "aab", expect: true},
		{s: "aa", t: "abaa", expect: true},
		{s: "aa", t: "abaacd", expect: true},
		{s: "", t: "abaacd", expect: true},
		{s: "", t: "d", expect: true},
	}

	for _, test := range testdata {
		get := canConstruct03(test.s, test.t)
		if get != test.expect {
			t.Fatalf("s:%v, t:%v, expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}

}
