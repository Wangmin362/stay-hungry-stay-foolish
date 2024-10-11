package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(claim(3, 2))
}

// 题目分析：完全背包+排列问题
// 遍历顺序： 先背包,在物品， 物品从小到大
func claim(n, m int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 0; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}

	return dp[n]
}
