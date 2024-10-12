package _0_basic

import (
	"fmt"
	"testing"
)

// 先考虑左边，在考虑右边
func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}

	cand := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		cand[i] = 1 // 每个小孩都需要分一颗糖果
	}

	// 先考虑左边的孩子
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && cand[i] <= cand[i-1] {
			cand[i] = cand[i-1] + 1
		}
	}

	// 在考虑右边的孩子
	res := cand[len(cand)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && cand[i] <= cand[i+1] {
			cand[i] = cand[i+1] + 1
		}
		res += cand[i]
	}

	return res
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
}
