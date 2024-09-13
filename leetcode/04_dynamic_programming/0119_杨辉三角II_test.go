package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/pascals-triangle-ii/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=EASY

// 对其之后可以发现规律
// [1]
// [1,1]
// [1,2,1]
// [1,3,3,1]
// [1,4,6,4,1]

func generateII(numRows int) []int {
	var prev []int
	for i := 0; i <= numRows; i++ {
		curr := make([]int, i+1)
		curr[0], curr[i] = 1, 1
		for j := 1; j < i; j++ {
			curr[j] = prev[j] + prev[j-1]
		}
		prev = curr
	}
	return prev
}

func TestGenerateII(t *testing.T) {
	res := generateII(3)
	fmt.Println(res)

}
