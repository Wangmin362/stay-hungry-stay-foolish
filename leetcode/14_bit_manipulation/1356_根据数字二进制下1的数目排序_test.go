package _0_basic

import (
	"fmt"
	"math/bits"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/

type binArr []int

func (b binArr) Len() int { return len(b) }

func (b binArr) Less(i, j int) bool {
	if bits.OnesCount(uint(b[i])) == bits.OnesCount(uint(b[j])) {
		return b[i] < b[j]
	}
	return bits.OnesCount(uint(b[i])) < bits.OnesCount(uint(b[j]))
}

func (b binArr) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func sortByBits(arr []int) []int {
	sort.Sort(binArr(arr))
	return arr
}

func TestSortByBits(t *testing.T) {
	//fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}
