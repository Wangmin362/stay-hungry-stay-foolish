package _0_basic

import "testing"

// 贪心思想：如果后一个属比前一个数字大，就在小数字买入，大数字卖出
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func TestMaxProfit(t *testing.T) {

}
