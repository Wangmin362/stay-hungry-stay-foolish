package _0_basic

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/group-anagrams/

// 解题思路：
// 解法一：把所有的字母对应的编码相加，只有出现相同次数的字符串才可能值一样。其实有可能40个字符串和50个字符的值是相等的，因此必须要求长度是相等的

func groupAnagrams(strs []string) [][]string {
	angMap := make(map[['z' - 'a' + 1]int][]string)

	for _, str := range strs {
		key := ['z' - 'a' + 1]int{}
		for _, sChar := range str {
			key[sChar-'a'] += 1
		}

		angMap[key] = append(angMap[key], str)
	}

	res := make([][]string, len(angMap))
	idx := 0
	for _, gr := range angMap {
		res[idx] = gr
		idx++
	}

	return res
}

func groupAnagrams01(strs []string) [][]string {
	cntFre := func(str string) [26]int {
		cnt := [26]int{}
		for idx := range str {
			cnt[str[idx]-'a']++
		}
		return cnt
	}

	resRaw := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := cntFre(str)
		arr, ok := resRaw[cnt]
		if ok {
			resRaw[cnt] = append(arr, str)
		} else {
			resRaw[cnt] = []string{str}
		}
	}

	var res [][]string
	for _, v := range resRaw {
		res = append(res, v)
	}
	return res
}

func groupAnagrams03(strs []string) [][]string {
	cache := make(map[[26]byte][]string)
	for i := 0; i < len(strs); i++ {
		var key [26]byte
		for j := 0; j < len(strs[i]); j++ {
			key[strs[i][j]-'a']++
		}
		v, ok := cache[key]
		if !ok {
			cache[key] = []string{strs[i]}
		} else {
			v = append(v, strs[i])
			cache[key] = v
		}
	}
	res := make([][]string, 0, len(cache))
	for _, v := range cache {
		res = append(res, v)
	}

	return res
}

func TestGroupAnagrams(t *testing.T) {
	var testdata = []struct {
		strs   []string
		expect [][]string
	}{
		//{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		//	expect: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{strs: []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			expect: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}}},
		{strs: []string{"duh", "ill"},
			expect: nil},
	}

	for _, test := range testdata {
		get := groupAnagrams03(test.strs)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("strs:%s, expect:%v, get:%v", test.strs, test.expect, get)
		}
	}

}
