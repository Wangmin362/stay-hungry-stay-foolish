package huawei_interview

import "testing"

// 题目：在一个国家仅有1分，2分，3分硬币，将钱N兑换成硬币有很多种兑法。请你编程序计算出共有多少种兑法。
// 输入：2934  输出：718831
// 输入：12553 输出：13137761

// 题目分析：这是一个完全背包问题，根据提议猜测是组合问题，不是排列问题，背包的容量为N
// 状态定义：dp[j]表示使用前i个物品装满容量为j的背包可能的方法数
// 地推公式：dp[j] += dp[j-coins[i]]
// 初始化：dp[0]=0
func changeMoney(coins []int, num int) int {
	dp := make([]int, num+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= num; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[num]
}

func TestChangeMoeny(t *testing.T) {
	var testdata = []struct {
		coins []int
		num   int
		want  int
	}{
		{coins: []int{1, 2, 3}, num: 2934, want: 718831},
		{coins: []int{1, 2, 3}, num: 12553, want: 13137761},
	}
	for _, tt := range testdata {
		get := changeMoney(tt.coins, tt.num)
		if get != tt.want {
			t.Fatalf("coins:%v, num:%v, want:%v, get:%v", tt.coins, tt.num, tt.want, get)
		}
	}
}
