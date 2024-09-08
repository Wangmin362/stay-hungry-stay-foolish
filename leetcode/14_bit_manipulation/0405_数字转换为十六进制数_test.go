package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/convert-a-number-to-hexadecimal/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func toHex(num int) string {
	mapping := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9',
		10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f',
	}
	if num == 0 {
		return "0"
	}

	abs := uint32(num)

	var res []byte
	for abs > 0 {
		mod := int(abs % 16)
		res = append(res, mapping[mod])
		abs >>= 4
	}

	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return string(res)
}

func TestToHex(t *testing.T) {
	fmt.Println(toHex(-1))
}
