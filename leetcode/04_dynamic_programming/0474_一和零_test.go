package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/ones-and-zeroes/description/

// 这道题目还是可以抽象为01背包问题。本质上就是背包的容量变成了一个二维数组，同事要满足m个0和n个1
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}

	cnt01 := func(str string) (int, int) {
		cnt0, cnt1 := 0, 0
		for _, c := range str {
			if c == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
		return cnt0, cnt1
	}

	// dp[m][n]定义为容量为m个0和n个1的背包，最多可以容量的字符串的数量
	// dp[m][n]的状态转移方程为 1、放入当前元素  2、不放入当前元素
	// dp[m][n] = max(dp[m][n], dp[m - len(0)][n - len(1)] + 1)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		cnt0, cnt1 := cnt01(str)
		for mm := m; mm >= cnt0; mm-- {
			for nn := n; nn >= cnt1; nn-- {
				noput := dp[mm][nn]
				put := dp[mm-cnt0][nn-cnt1] + 1
				dp[mm][nn] = int(math.Max(float64(noput), float64(put)))
			}
		}
	}

	return dp[m][n]
}

func findMaxForm02(strs []string, m int, n int) int {
	// dp[i][j]定义为i个0，j个1最多的结合个数
	// dp[i][j] = max(dp[i][j], dp[i-x][j-y]+1)
	get01 := func(str string) (int, int) {
		cnt0, cnt1 := 0, 0
		for _, c := range str {
			if c == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
		return cnt0, cnt1
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		x, y := get01(str)
		for i := m; i >= x; i-- {
			for j := n; j >= y; j-- {
				dp[i][j] = max(dp[i][j], dp[i-x][j-y]+1)
			}
		}
	}

	return dp[m][n]
}

func TestFindMaxForm(t *testing.T) {
	fmt.Println(findMaxForm02([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
