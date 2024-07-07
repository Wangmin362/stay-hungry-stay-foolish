package _0_basic

import (
	"math"
)

// 教程：https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE
// 题目：https://kamacoder.com/problempage.php?pid=1046

// 二维数组
func getMacValue01(space int, weight, value []int) int {
	// dp[i][j]为0..i个物品，放进容量为j的背包的最大价值
	// dp[i][j] = max(dp[i-1][j] ,     dp[i-1][j - weight[i]] + value[i])
	//                 不放物品i的价值       放入物品i的价值
	// 为什么不妨物品i的价值就是dp[i-1][j]，因为如果背包不妨物品i，那么背包的价值就是物品1..i-1所有物品的价值之和，此时物品i对于背包剩余空间
	// 没有影响。
	// 为什么背包放入物品i的公式为：dp[i-1][j-weight[i]] + value[i]，应为如果物品i想要放入到背包中，那么放入背包前，背包的容量一定是
	// [j-weight[i]]或者比这个更大，应为只有这样，物品i才能放入到背包容量为[j-weight[i]]的物品，最后在加上物品i的价值即可

	dp := make([][]int, len(weight))
	dp[0] = make([]int, space+1)
	for j := 0; j <= space; j++ { // 初始化第一行
		if weight[0] > j { // 如果物品的重量大于剩余空间，那么背包什么也放不进去，初始化为0
			dp[0][j] = 0
		} else { // 否则，设置物品放得进去，就初始化为物品0的价值
			dp[0][j] = value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		dp[i] = make([]int, space+1)
		dp[i][0] = 0 // 背包的容量为0，那么最大价值只能为0
		for j := 1; j <= space; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j] // 当前背包已经放不进去物品i了，最大价值只能是前[1, i-1]个物品价值之和
			} else {
				notPut := dp[i-1][j]                   // 不放入当前物品
				put := dp[i-1][j-weight[i]] + value[i] // 放入物品i
				dp[i][j] = int(math.Max(float64(notPut), float64(put)))
			}
		}
	}

	return dp[len(weight)-1][space+1]
}

// 一维数组，亦称之为滚动数组，为什么能压缩到一维数组呢？ 其实从上面的二维数组递推公式我们可以发现，每一行的数据只和前一行的数据有关，和其他行无关
// 也就是说当前行的数据其实是根据上一行的数据加工出来的
func getMacValue(space int, weight, value []int) int {
	// dp[j]为容量为j的背包，可以容量物品的最大价值，每一行的意义是不一样的，因为每一行是在前一行的基础之上增加了当前的物品
	// dp[j] = max(dp[j], dp[j-weight(i)] + value[i])，其实就是上一行容量为j的背包的价值，和放入物品i的价值，二者最大值
	dp := make([]int, space+1)
	for j := 0; j <= space; j++ { // 使用物品0来初始化第一行数据
		if weight[0] > space {
			dp[j] = 0
		} else {
			dp[j] = value[0]
		}
	}
	for i := 0; i < len(weight); i++ {
		for j := space; j >= weight[i]; j-- {
			notPut := dp[j]
			put := dp[j-weight[i]] + value[i]
			dp[j] = int(math.Max(float64(notPut), float64(put)))
		}
	}

	return dp[space]
}
