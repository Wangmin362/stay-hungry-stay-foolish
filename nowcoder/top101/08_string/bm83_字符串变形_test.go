package _8_string

import (
	"fmt"
	"strings"
	"testing"
)

func reverseString(str string) string {
	bStr := []byte(str)
	diff := byte('a' - 'A')
	for idx := range bStr {
		if bStr[idx] >= 'A' && bStr[idx] <= 'Z' {
			bStr[idx] += diff
		} else if bStr[idx] >= 'a' && bStr[idx] <= 'z' {
			bStr[idx] -= diff
		}
	}
	return string(bStr)
}

func trans(s string, n int) string {
	sp := strings.Split(s, " ")
	var res []string
	for i := len(sp) - 1; i >= 0; i-- {
		ss := reverseString(sp[i])
		res = append(res, ss)
	}

	return strings.Join(res, " ")
}

func TestTrans(t *testing.T) {
	fmt.Println(trans("This is a sample", 5))
}
