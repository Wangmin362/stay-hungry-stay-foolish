package _8_string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func solve(s string, t string) string {
	reverse := func(str []string) {
		left, right := 0, len(str)-1
		for left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}

	sps := strings.Split(s, "")
	reverse(sps)
	spt := strings.Split(t, "")
	reverse(spt)
	var res []string
	idx := 0
	overlfow := 0
	for idx < len(sps) || idx < len(spt) {
		sNum, tNum := 0, 0
		if idx < len(sps) {
			sNum, _ = strconv.Atoi(sps[idx])
		}
		if idx < len(spt) {
			tNum, _ = strconv.Atoi(spt[idx])
		}
		sum := sNum + tNum + overlfow
		if sum > 9 {
			overlfow = 1
			sum %= 10
		} else {
			overlfow = 0
		}
		res = append(res, strconv.Itoa(sum))
		idx++
	}
	if overlfow == 1 {
		res = append(res, "1")
	}

	reverse(res)
	return strings.Join(res, "")
}

func TestSolve(t *testing.T) {
	fmt.Println(solve("99999", "1"))
}
