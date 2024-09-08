package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/binary-gap/

func binaryGap(n int) int {
	preOneIdx := -1
	idx := 0
	maxDist := 0
	for n > 0 {
		if n&1 == 1 {
			if preOneIdx == -1 {
				preOneIdx = idx
			} else {
				if maxDist < idx-preOneIdx {
					maxDist = idx - preOneIdx
				}
				preOneIdx = idx
			}
		}
		idx++
		n >>= 1
	}
	return maxDist
}

func TestBinaryGap(t *testing.T) {
	fmt.Println(binaryGap(22))
}
