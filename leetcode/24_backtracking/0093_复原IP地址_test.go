package _1_array

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/restore-ip-addresses/description/

func restoreIpAddresses(s string) []string {
	var backtracking func(s string, startIdx int)

	validMap := map[string]bool{}
	isValidIpNum := func(str string) bool {
		if valid, ok := validMap[str]; ok {
			return valid
		}

		if len(str) > 1 && str[0] == '0' { // 任何前导0都无效
			validMap[str] = false
			return false
		}
		if strings.Compare(fmt.Sprintf("%3s", str), "255") <= 0 {
			validMap[str] = true
			return true
		}

		validMap[str] = false
		return false
	}

	var res []string
	var path []string
	backtracking = func(s string, startIdx int) {
		if startIdx >= len(s) {
			if len(path) == 4 {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		for idx := startIdx; idx < len(s) && idx <= startIdx+2; idx++ {
			if len(path) == 4 && startIdx < len(s) {
				continue
			}
			str := s[startIdx : idx+1]
			if !isValidIpNum(str) {
				continue
			}
			path = append(path, str)
			backtracking(s, idx+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0)
	return res
}

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("101023"))
}
