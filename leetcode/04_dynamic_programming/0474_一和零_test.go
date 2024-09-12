package _0_basic

import (
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

// 直接使用回溯解决
func findMaxFormBacktrackint(strs []string, m int, n int) int {
	var backtracking func(strat int, mx, nx int)

	cnt01 := func(str string) (int, int) {
		zero, one := 0, 0
		for i := range str {
			if str[i] == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}

	var res int
	var path []string
	backtracking = func(strat int, mx, nx int) {
		if mx <= m && nx <= n {
			res = max(res, len(path))
		}
		if mx > m || nx > n {
			return
		}

		for i := strat; i < len(strs); i++ {
			zero, one := cnt01(strs[i])
			path = append(path, strs[i])
			backtracking(i+1, mx+zero, nx+one)
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0, 0)
	return res
}

// 题目分析：抽象为背包的0容量为m, 1容量为n的背包，最多可以装多少个字符串
// 明确定义：dp[m][n]定义为背包为m个0， n个1的背包，最多可以装字符串的数量，这个背包的容量是2维的
// 状态转移方程：dp[m][n] = max(dp[m-zero(str[i])][z-one(str[i]) + 1], dp[m][n])
// 初始化：dp[0][0] = 0
// 遍历顺序：先物品，再背包容量，容量需要倒叙，防止一个物品放入多次
// dp数组大小 [m+1][n+1]
// 返回值：dp[m][n]
func findMaxForm0912(strs []string, m int, n int) int {
	cnt01 := func(str string) (int, int) {
		zero, one := 0, 0
		for i := range str {
			if str[i] == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 0; i < len(strs); i++ {
		zero, one := cnt01(strs[i])
		for mi := m; mi >= zero; mi-- {
			for ni := n; ni >= one; ni-- {
				dp[mi][ni] = max(dp[mi-zero][ni-one]+1, dp[mi][ni])
			}
		}
	}

	return dp[m][n]
}

func TestFindMaxForm(t *testing.T) {
	var testdata = []struct {
		strs []string
		m    int
		n    int
		want int
	}{
		{strs: []string{"10", "0001", "111001", "1", "0"}, m: 5, n: 3, want: 4},
	}
	for _, tt := range testdata {
		get := findMaxForm0912(tt.strs, tt.m, tt.n)
		if get != tt.want {
			t.Fatalf("strs:%v, m:%v, n:%v want:%v, get:%v", tt.strs, tt.m, tt.n, tt.want, get)
		}
	}
}
