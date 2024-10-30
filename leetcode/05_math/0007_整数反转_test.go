package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) int {
	if x == 0 {
		return x
	}

	isPos := true
	if x < 0 {
		isPos = false
		x = -x
	}

	var res int
	for x > 0 {
		mod := x % 10
		x /= 10
		if res*10 > math.MaxInt32 || res*10+mod > math.MaxInt32 {
			return 0
		}
		res = res*10 + mod
	}

	if !isPos {
		return -res
	}

	return res
}

func TestReverse(t *testing.T) {
	fmt.Println(reverse(1210))
	fmt.Println(reverse(-457))
	fmt.Println(reverse(1534236469))
}
