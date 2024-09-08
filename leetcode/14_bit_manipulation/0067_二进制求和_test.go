package _0_basic

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/add-binary/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func addBinary(a string, b string) string {

	revrese := func(arr []byte) {
		left, right := 0, len(arr)-1
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	aa := []byte(a)
	revrese(aa)
	bb := []byte(b)
	revrese(bb)

	var res []string
	overflow := 0
	for idx := 0; idx < len(aa) || idx < len(bb); idx++ {
		an, bn := 0, 0
		if idx < len(aa) {
			an = int(aa[idx] - '0')
		}
		if idx < len(bb) {
			bn = int(bb[idx] - '0')
		}
		sum := an + bn + overflow
		if sum > 1 {
			overflow = 1
			sum %= 2
		} else {
			overflow = 0
		}
		res = append(res, strconv.Itoa(sum))
	}
	if overflow == 1 {
		res = append(res, "1")
	}

	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return strings.Join(res, "")
}

func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("110", "10"))
}
