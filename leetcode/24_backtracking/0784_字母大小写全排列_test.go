package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/letter-case-permutation/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func letterCasePermutation(s string) []string {
	var backtracking func(idx int)

	sp := []byte(s)
	var res []string
	choices := [2]bool{true, false} // 变和不变
	backtracking = func(idx int) {
		// 数字直接添加
		for idx < len(sp) && sp[idx] >= '0' && sp[idx] <= '9' {
			idx++
		}

		if idx >= len(s) {
			res = append(res, string(sp))
			return
		}

		// 英文字母才需要处理变化还是不变
		for _, choice := range choices {
			c := sp[idx]

			if choice {
				if c >= 'a' && c <= 'z' { // 修改为大写字母
					sp[idx] -= 'a' - 'A'
					backtracking(idx + 1)
					sp[idx] += 'a' - 'A'
				} else { // 修改为小写字母
					sp[idx] += 'a' - 'A'
					backtracking(idx + 1)
					sp[idx] -= 'a' - 'A'
				}

			} else {
				backtracking(idx + 1)
			}
		}
	}

	backtracking(0)
	return res
}

func TestLetterCasePermutation(t *testing.T) {
	fmt.Println(letterCasePermutation("a1b2"))
}
