package _1_array

import (
	"slices"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	isAlpha := func(idx int) bool {
		if paragraph[idx] >= 'a' && paragraph[idx] <= 'z' {
			return true
		}
		return false
	}
	paragraph = strings.ToLower(paragraph)
	mapping := make(map[string]int)
	idx := 0
	first := -1
	for idx < len(paragraph) {
		for idx < len(paragraph) && !isAlpha(idx) { // 找到单词第一个字符
			idx++
		}
		if idx < len(paragraph) {
			first = idx
		}
		for idx < len(paragraph) && isAlpha(idx) { // 找到单词最后一个字符
			idx++
		}
		if first != -1 {
			str := paragraph[first:idx]
			if !slices.Contains(banned, str) {
				mapping[str]++
			}
		}
		idx++
	}

	var maxFrq int
	var maxFrqStr string
	for str, frq := range mapping {
		if frq > maxFrq {
			maxFrq = frq
			maxFrqStr = str
		}
	}

	return maxFrqStr

}
