package _1_array

import (
	"fmt"
	"testing"
)

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	if len(s) < minSize {
		return 0
	}

	cm := make(map[byte]int)          // 统计窗口内字符的频率
	windowStr := make(map[string]int) // 统计符合条件的窗口字符串的频率
	left, right := 0, 0
	for ; right < len(s); right++ {
		cm[s[right]]++
		if right-left+1 < maxSize { // 窗口不够最大窗口
			if right-left+1 >= minSize && len(cm) <= maxLetters {
				windowStr[s[left:right+1]]++
			}
			continue
		}
		if right-left+1 >= minSize && len(cm) <= maxLetters {
			windowStr[s[left:right+1]]++
		}

		for right-left+1 >= minSize {
			cm[s[left]]--
			if cm[s[left]] == 0 {
				delete(cm, s[left])
			}
			left++

			if right-left+1 >= minSize && len(cm) <= maxLetters {
				windowStr[s[left:right+1]]++
			}
		}
	}
	right--

	for right-left+1 >= minSize {
		cm[s[left]]--
		if cm[s[left]] == 0 {
			delete(cm, s[left])
		}
		left++

		if right-left+1 >= minSize && len(cm) <= maxLetters {
			windowStr[s[left:right+1]]++
		}
	}

	maxCnt := 0
	for str, cnt := range windowStr {
		fmt.Printf("%v -> %v\n", str, cnt)
		maxCnt = max(maxCnt, cnt)
	}

	return maxCnt
}

func TestMaxFreq(t *testing.T) {
	var testdata = []struct {
		s          string
		maxLetters int
		minSize    int
		maxSize    int
		want       int
	}{
		//{s: "aababcaab", maxLetters: 2, minSize: 3, maxSize: 4, want: 2},
		//{s: "aaaa", maxLetters: 1, minSize: 3, maxSize: 3, want: 2},
		//{s: "aabcabcab", maxLetters: 2, minSize: 2, maxSize: 3, want: 3},
		//{s: "abcde", maxLetters: 2, minSize: 3, maxSize: 3, want: 0},
		{s: "abcabababacabcabc", maxLetters: 3, minSize: 3, maxSize: 10, want: 2},
	}
	for _, tt := range testdata {
		get := maxFreq(tt.s, tt.maxLetters, tt.minSize, tt.maxSize)
		if get != tt.want {
			t.Errorf("s:%v, maxLetters:%v, minSize:%v, maxSize:%v, want:%v, get:%v",
				tt.s, tt.maxLetters, tt.minSize, tt.maxSize, tt.want, get)
		}
	}
}
