package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	mn := strings.Split(scan.Text(), " ")
	m, _ := strconv.Atoi(mn[0])
	capacity, _ := strconv.Atoi(mn[1])

	scan.Scan()
	ws := strings.Split(scan.Text(), " ")
	weights := make([]int, m)
	for idx, w := range ws {
		weights[idx], _ = strconv.Atoi(w)
	}

	scan.Scan()
	vs := strings.Split(scan.Text(), " ")
	values := make([]int, m)
	for idx, v := range vs {
		values[idx], _ = strconv.Atoi(v)
	}

	val := maxValue02(capacity, weights, values)
	fmt.Println(val)
}

// 明确定义：dp[i][j]容量为j的背包放入前i个物品的最大价值
// 递推公式：dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i]] + value[i])
// 初始化：初始化第一行，初始化第一列
// 遍历顺序：先物品，后背包，从小到大
func maxValue01(capacity int, weights, values []int) int {
	dp := make([][]int, len(weights))
	for i := 0; i < len(weights); i++ {
		dp[i] = make([]int, capacity+1)
	}
	// 初始化第一行
	for j := weights[0]; j <= capacity; j++ {
		dp[0][j] = values[0]
	}
	// 第一列，容量为0， 自然做大价值肯定是0，不需要初始化

	for i := 1; i < len(weights); i++ {
		for j := 0; j <= capacity; j++ {
			if j >= weights[i] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i]]+values[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(weights)-1][capacity]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 状态压缩
// 明确定义：dp[j]表示前i个物品放入容量为j的背包的最大价值
// 分析： 二维dp的递推公式为：dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i]] + value[i]), 可以看出当前状态之和上一行的状态和本行的状态有关
// 递推公式：dp[j] = max(dp[j], dp[j-weight[i]] + values[i])
// 初始化：初始化第一个物品
// 遍历顺序：先物品，后容量。 物品从小到大，容量从大到小（如果是多重背包，容量也应该从小到大）
func maxValue02(capacity int, weights, values []int) int {
	dp := make([]int, capacity+1)
	for j := weights[0]; j <= capacity; j++ {
		dp[j] = values[0]
	}
	for i := 1; i < len(weights); i++ {
		for j := capacity; j >= 0; j-- {
			if j >= weights[i] {
				dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
			}
		}
	}
	return dp[capacity]
}
