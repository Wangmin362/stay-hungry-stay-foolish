package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

func maxVowels(s string, k int) int {
	isVowel := func(char uint8) bool {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			return true
		}
		return false
	}

	k = min(len(s), k)
	maxLen := 0
	currLen := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currLen++
		}
	}
	maxLen = max(maxLen, currLen)
	left, right := 0, k-1
	for right < len(s) {
		right++
		if right < len(s) && isVowel(s[right]) {
			currLen++
		}
		if isVowel(s[left]) {
			currLen--
		}
		left++
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}

func maxVowels0901(s string, k int) int {
	isVowels := func(c byte) bool {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return true
		}
		return false
	}

	var maxVol int
	var fast int
	for fast = 0; fast < k && fast < len(s); fast++ {
		if isVowels(s[fast]) {
			maxVol++
		}
	}
	if len(s) <= k {
		return maxVol
	}
	slow := 1
	cnt := maxVol
	for fast < len(s) {
		if isVowels(s[slow-1]) {
			cnt--
		}
		if isVowels(s[fast]) {
			cnt++
		}
		fast++
		slow++
		if cnt > maxVol {
			maxVol = cnt
		}
	}

	return maxVol
}

// 模板
func maxVowels0901_2(s string, k int) int {
	var ans int
	var cur int
	for idx, in := range s {
		// 判断窗口右边界
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			cur++
		}
		if idx < k-1 {
			// 可能存在字符串长度不足k个，虽然题目已经保证
			ans = max(cur, ans)
			continue // 窗口不足时，不需要判断左边界
		}

		// 更新最大值
		ans = max(cur, ans)

		// 判断左边界
		out := s[idx-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			cur--
		}
	}
	return ans
}

func TestMaxVowels(t *testing.T) {
	testdata := []struct {
		s      string
		t      int
		expect int
	}{
		{s: "abciiidef", t: 3, expect: 3},
	}

	for _, test := range testdata {
		get := maxVowels0901_2(test.s, test.t)
		if get != test.expect {
			t.Errorf("s:%v, t:%v  expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}
}
