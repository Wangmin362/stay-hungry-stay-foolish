package _0_basic

import (
	"fmt"
	"testing"
)

// 教程：https://www.programmercarl.com/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF%E5%AE%8C%E5%85%A8%E8%83%8C%E5%8C%85%E7%89%88%E6%9C%AC.html#%E6%80%9D%E8%B7%AF
// 题目：https://kamacoder.com/problempage.php?pid=1067

// 题目抽象：背包的容量为n, 一共有m个物品，问m个物品填满背包有多少种方式
func scrapeFloor01(n, m int) int {
	// dp[j]定义为容量为j的背包，填满dp[j]的次数
	// dp[j] += dp[j - num[i]]
	// 结果为dp[n]
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 0; j <= n; j++ { // 遍历背包
		for i := 1; i <= m; i++ { // 遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
		fmt.Println("背包的容量为：", j, dp)
	}
	return dp[n]
}

func scrapeFloor02(n, m int) int {
	// dp[j]定义为容量为j的背包，填满dp[j]的次数
	// dp[j] += dp[j - num[i]]
	// 结果为dp[n]
	dp := make([]int, n+1)
	dp[0] = 1
	fmt.Println("物品：", 0, dp)
	for i := 1; i <= m; i++ { // 遍历物品
		for j := i; j <= n; j++ { // 遍历背包
			dp[j] += dp[j-i]
		}
		fmt.Println("物品：", i, dp)
	}
	return dp[n]
}

// 排列问题， 背包容量为n,物品为m, 并且是1到m. 由于是排列问题，因此是先背包，再物品，都是从小打到
// dp[j] += dp[j-m]

func scrapeFloor0912(n, m int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for j := 1; j <= n; j++ { // 背包
		for i := 1; i <= m && i <= j; i++ { // 物品
			dp[j] += dp[j-i]
		}
		fmt.Println(dp)
	}

	return dp[n]
}

func TestScrapeFloor(t *testing.T) {
	var testdata = []struct {
		n    int
		m    int
		want int
	}{
		{n: 3, m: 2, want: 3},
	}
	for _, tt := range testdata {
		get := scrapeFloor0912(tt.n, tt.m)
		if get != tt.want {
			t.Fatalf("n:%v, m:%v, want:%v, get:%v", tt.n, tt.m, tt.want, get)
		}
	}
}
