package _1_array

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/?envType=study-plan-v2&envId=top-interview-150

func findSubstring(s string, words []string) []int {
	sort.Strings(words) // 单词排序，避免重复组合

	var backtracking func()
	var allPtn []string
	used := make(map[int]bool)
	var path []string
	backtracking = func() {
		if len(path) == len(words) {
			allPtn = append(allPtn, strings.Join(path, ""))
			return
		}
		for i := 0; i < len(words); i++ {
			if used[i] || (i > 0 && words[i] == words[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, words[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtracking()
	fmt.Println(allPtn)

	var res []int
	for i := 0; i < len(allPtn); i++ {
		cnt := 0
		for j := 0; j < len(s); {
			str := allPtn[i]
			ns := s[j:]
			idx := strings.Index(ns, str)
			if idx >= 0 { // 说明找到了
				res = append(res, idx+cnt*len(words[i]))
				cnt++
				j = idx + len(words[i]) // s中可能包含多个word[i]
			} else {
				break
			}
		}
	}

	return res
}

func TestFindSubstring(t *testing.T) {
	var teatdata = []struct {
		s     string
		words []string
		want  []int
	}{
		//{s: "barfoothefoobarman", words: []string{"foo", "bar"}, want: []int{0, 9}},
		{s: "foobarfoobar", words: []string{"foo", "bar"}, want: []int{0, 3, 6}},
	}

	for _, test := range teatdata {
		sum01 := findSubstring(test.s, test.words)
		if !reflect.DeepEqual(sum01, test.want) {
			t.Errorf("s:%v, words:%v, want:%v, get:%v", test.s, test.words, test.want, sum01)
		}
	}
}
