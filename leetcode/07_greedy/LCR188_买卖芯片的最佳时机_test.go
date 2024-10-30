package _0_basic

import (
	"fmt"
	"testing"
)

func bestTiming(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res, minPrice := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		res = max(res, prices[i]-minPrice)
	}
	return res
}

func TestBestTiming(t *testing.T) {
	fmt.Println(bestTiming([]int{3, 6, 2, 9, 8, 5}))
}
