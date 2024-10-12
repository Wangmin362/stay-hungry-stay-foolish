package _0_basic

import (
	"fmt"
	"testing"
)

// 贪心思路：把第一个字符放在区间当中，像后面找，找到最后一次出现这个字符的位置，与此同时，期间遍历过的所有字符，都需要找最后一次
// TODO 这种思路似乎有问题
func partitionLabels(s string) []int {
	if len(s) < 2 {
		return []int{len(s)}
	}

	// 下标0表示第一次出现的位置，下标二表示最后一次出现的位置
	cache := [26][2]int{}
	for i := 0; i < 26; i++ {
		cache[i] = [2]int{-1, -1} // 如果是-1,说明这个字符没有出现，不需要考虑
	}

	for i := 0; i < len(s); i++ {
		if cache[s[i]-'a'][0] == -1 { // 第一次找到这个字符
			cache[s[i]-'a'][0] = i
			cache[s[i]-'a'][1] = i // 可能这个字符就出现过一次，所以第一次要更新为第一次出现的位置
		} else {
			cache[s[i]-'a'][1] = i // 更新这个字符的最后一次位置
		}
	}

	part := make([][2]int, 0, len(s))
	for i := 0; i < len(cache); i++ {
		if cache[i][0] != -1 {
			part = append(part, cache[i])
		}
	}

	var res []int
	var lastIdx int
	for i := 1; i < len(part); i++ {
		if part[i][0] < part[i-1][1] { // 说明两个区间重叠
			part[i][1] = max(part[i][1], part[i-1][1]) // 那么区间，有边界取最大值
		} else { // 没有重叠
			res = append(res, part[i][0]-lastIdx)
			lastIdx = i
		}
	}

	return res
}

//func partitionLabels02(s string) []int {
//	if len(s) < 2 {
//		return []int{len(s)}
//	}
//
//	c := make(map[byte]int, len(s))
//	for i := 0; i < len(s); i++ {
//		c[s[i]] = i // 记录每个字符最后一次出现的位置
//	}
//
//	idxs := make([]int, len(s))
//	for i := 0; i < len(s); i++ {
//		idxs[i] = c[s[i]]
//	}
//
//	var res []int
//	maxBorder := idxs[0]
//	for i := 1; i < len(idxs); i++ {
//		if idxs[i] > maxBorder {
//
//		}
//	}
//}

func TestPartitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
