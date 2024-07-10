package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// 教程：https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE
// 题目：https://kamacoder.com/problempage.php?pid=1046

// 二维数组
/*
         0 1 2 3 4 5 6 7 8  背包容量
物品： 0 [0 1 1 1 1 1 1 1 1]
物品： 1 [0 1 1 3 4 4 4 4 4]
物品： 2 [0 1 1 3 4 5 6 6 8]
物品： 3 [0 1 1 3 4 5 6 7 8]

*/
func getMaxValue01(space int, weight, value []int) int {
	// dp[i][j]为0..i个物品，放进容量为j的背包的最大价值
	// dp[i][j] = max(dp[i-1][j] ,     dp[i-1][j - weight[i]] + value[i])
	//                 不放物品i的价值       放入物品i的价值
	// 为什么不妨物品i的价值就是dp[i-1][j]，因为如果背包不妨物品i，那么背包的价值就是物品1..i-1所有物品的价值之和，此时物品i对于背包剩余空间
	// 没有影响。
	// 为什么背包放入物品i的公式为：dp[i-1][j-weight[i]] + value[i]，应为如果物品i想要放入到背包中，那么放入背包前，背包的容量一定是
	// [j-weight[i]]或者比这个更大，应为只有这样，物品i才能放入到背包容量为[j-weight[i]]的物品，最后在加上物品i的价值即可

	dp := make([][]int, len(weight))
	dp[0] = make([]int, space+1)

	// 1、初始化第一行  如果只有第零个物品，那么不管背包容量为多大，只要背包的容量大于物品的重量，背包就能放入这个物品，此时背包的最大价值只可能是第零个物品的价值
	// Q、为什么要先初始化第0行的数据呢？
	// A、原因是因为从公式dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + value[i])中可以知道每一行是依赖前一行的，因此我们初始化了第0
	// 行的数据。 相当于问题是当背包的容量为0..j时，放入物品0的最大价值是多杀
	for j := weight[0]; j <= space; j++ {
		dp[0][j] = value[0]
	}
	fmt.Println("物品：", 0, dp[0])
	for i := 1; i < len(weight); i++ { // 从第一个物品开始计算计算，因为第0行已经初始化了
		dp[i] = make([]int, space+1)  // 分配空间，默认初始化为0
		for j := 1; j <= space; j++ { // 从背包容量为1开始计算，因为背包容量为0时，不管有多少物品，肯定都放不进去，因此背包的最大价值为0
			if weight[i] > j {
				dp[i][j] = dp[i-1][j] // 当前背包已经放不进去物品i了，最大价值只能是前[1, i-1]个物品价值之和
			} else {
				notPut := dp[i-1][j]                   // 不放入当前物品
				put := dp[i-1][j-weight[i]] + value[i] // 放入物品i
				dp[i][j] = int(math.Max(float64(notPut), float64(put)))
			}
		}
		fmt.Println("物品：", i, dp[i])
	}

	return dp[len(weight)-1][space]
}

// 一维数组，亦称之为滚动数组，为什么能压缩到一维数组呢？ 其实从上面的二维数组递推公式我们可以发现，每一行的数据只和前一行的数据有关，和其他行无关
// 也就是说当前行的数据其实是根据上一行的数据加工出来的
func getMaxValue02(space int, weight, value []int) int {
	// dp[j]为容量为j的背包，可以容量物品的最大价值，每一行的意义是不一样的，因为每一行是在前一行的基础之上增加了当前的物品
	// dp[j] = max(dp[j], dp[j-weight(i)] + value[i])，其实就是上一行容量为j的背包的价值，和放入物品i的价值，二者最大值
	dp := make([]int, space+1)
	for j := weight[0]; j <= space; j++ { // 使用物品0来初始化第一行数据
		dp[j] = value[0]
	}
	fmt.Println("物品：", 0, dp)
	for i := 1; i < len(weight); i++ {
		for j := space; j >= weight[i]; j-- {
			notPut := dp[j]
			put := dp[j-weight[i]] + value[i]
			dp[j] = int(math.Max(float64(notPut), float64(put)))
		}
		fmt.Println("物品：", 0, dp)
	}

	return dp[space]
}

func TestGetMaxValue(t *testing.T) {
	getMaxValue02(8, []int{1, 3, 5, 7}, []int{1, 3, 5, 7})
}
