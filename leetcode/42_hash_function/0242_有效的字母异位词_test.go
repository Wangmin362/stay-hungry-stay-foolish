package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/valid-anagram/

// 解题思路：
// 解法一：使用hashmap存储s字符串，然后遍历t字符串，如果map中没有，则说明肯定不是相似的字符串
// 解法二：可以使用数组代替字符串，因为一共只有26个英文字母，但是如果是Unicode，只能使用Map,因为无法确定数组的大小

func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // 长度不等，肯定不是
		return false
	}

	sMap := make(map[rune]int, len(s))
	for _, sChar := range s {
		cnt, ok := sMap[sChar]
		if !ok {
			sMap[sChar] = 1
		} else {
			sMap[sChar] = cnt + 1
		}
	}

	for _, tChar := range t {
		cnt, ok := sMap[tChar]
		if !ok {
			return false
		} else {
			if cnt-1 == 0 {
				delete(sMap, tChar)
			} else {
				sMap[tChar] = cnt - 1
			}
		}
	}

	if len(sMap) != 0 {
		return false
	}

	return true
}

// map
func isAnagram01(s string, t string) bool {
	freCnt := func(str string) map[uint8]int {
		m := make(map[uint8]int)
		for idx := range str {
			m[str[idx]]++
		}
		return m
	}

	sCnt := freCnt(s)
	tCnt := freCnt(t)
	return reflect.DeepEqual(sCnt, tCnt)
}

// 数组
func isAnagram02(s string, t string) bool {
	freCnt := func(str string) [26]int {
		cnt := [26]int{}
		for idx := range str {
			cnt[str[idx]-'a']++
		}
		return cnt
	}

	sCnt := freCnt(s)
	tCnt := freCnt(t)
	return reflect.DeepEqual(sCnt, tCnt)
}

func TestRemoveNthFromEnd(t *testing.T) {
	var testdata = []struct {
		s      string
		t      string
		expect bool
	}{
		{s: "abcd", t: "dbca", expect: true},
		{s: "", t: "", expect: true},
		{s: "d", t: "s", expect: false},
		{s: "ssssssssssss", t: "ssssssssssss", expect: true},
	}

	for _, test := range testdata {
		get := isAnagram02(test.s, test.t)
		if get != test.expect {
			t.Fatalf("s:%v, t:%v, expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}

}
