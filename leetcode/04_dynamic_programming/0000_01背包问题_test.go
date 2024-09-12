package _0_basic

import (
	"fmt"
	"testing"
)

// 教程：https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE
// 题目：https://kamacoder.com/problempage.php?pid=1046

// 题目：背包的容量为4， 有3个物品，每个物品只有一个，问背包可以装物品的最大值  物品的重量和价值为
// 		  重量   价值
// 物品0    1     15
// 物品1    3     20
// 物品2    4     30

// 题目分析：
// 解法一：每个物品可以放入到背包中，也可以不放入到背包当中，也就是要么放要么不妨。因此可以使用回溯的方式暴力搜索所有可能的结果，然后找到最大值即可

func BaggageMaxValueBacktracking(weight, value []int, cap int) int {
	var backtracking func(start, currWeight, currValue int)

	var res int
	backtracking = func(start, currWeight, currValue int) {
		res = max(res, currValue)

		for i := start; i < len(weight); i++ {
			if currWeight+weight[i] > cap { // 放不进去直接跳过
				continue
			}

			currWeight += weight[i]
			currValue += value[i]
			backtracking(i+1, currWeight, currValue)
			currValue -= value[i]
			currWeight -= weight[i]
		}
	}

	backtracking(0, 0, 0)
	return res
}

// 解法二：其实当前物品放与不妨取决于之前的状态，因此可以使用动态规划的方式来记住之前的状态，从而节省时间
// 明确定义：dp[i][j]表示前i个物品放入背包容量为j的背包的最大价值
// 状态方程：dp[i][j] = max(dp[i-1][j-weight[i]] + value[i], dp[i-1][j]) 即当前的最大价值，其实就是放入物品i的最大价值，和不放入物品
// i的最大值，若不放入物品i，那么背包的最大价值就是dp[i-1][j]，即前i-1个物品放入背包的最大价值。如果放入这个物品，显然背包要保留足够的空间，可以放入
// 物品i，因此需要找到一个容量为j-weight[i]的背包，这样才能够放入物品i，然后用前i-1个物品放入容量为j-weight[i]的背包。所以背包的不放入物品
// i的价值为dp[i-1][j-weight[i]] + value[i]
// 初始化：根据推导公式，dp[i][j]取决于前一行的状态，因此dp[0][j]必须要初始化，与此同时当容量为0的时候，最大价值一定为0，也就是说第一列初始化为0
// 遍历顺序：从前往后，从上往下，先遍历物品，在遍历容量  先遍历物品，在遍历容量
// dp数组大小：dp := make([][]int, len(weight)) dp[0]= make([]int, cap+1)  因为讨论的是容量0到cap，因此是cap+1
// 返回值 dp[len(weight)-1][cap]
// 		  重量   价值
// 物品0    1     15
// 物品1    3     20
// 物品2    4     30

// 手动模拟：
//
//	       容量0   1    2    3    4
//	物品0    0     15   15   15   15
//	物品1    0     15   15   20   35
//	物品2    0     15   15   20   35
func BaggageMaxValueDp01(weight, value []int, cap int) int {
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, cap+1)
	}

	// 初始化第一行，第一列不需要初始化，因为默认值就是0
	for j := weight[0]; j <= cap; j++ {
		dp[0][j] = value[0]
	}

	fmt.Println(dp[0])
	for i := 1; i < len(weight); i++ {
		for j := 1; j <= cap; j++ {
			if j < weight[i] { // 当前的背包容量根本放不进去物品i
				dp[i][j] = dp[i-1][j] //就只能不放入物品i
			} else {
				dp[i][j] = max(dp[i-1][j-weight[i]]+value[i], dp[i-1][j])
			}
		}
		fmt.Println(dp[1])
	}

	return dp[len(weight)-1][cap]
}

// 遍历顺序：先容量，在物品
// 		  重量   价值
// 物品0    1     15
// 物品1    3     20
// 物品2    4     30

// 手动模拟：
//
//	       容量0   1    2    3    4
//	物品0    0     15   15   15   15
//	物品1    0     15   15   20   35
//	物品2    0     15   15   20   35
func BaggageMaxValueDp02(weight, value []int, cap int) int {
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, cap+1)
	}

	for j := weight[0]; j <= cap; j++ {
		dp[0][j] = value[0]
	}

	for j := 1; j <= cap; j++ {
		for i := 1; i < len(weight); i++ {
			if j < weight[i] { // 放不进去
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-weight[i]]+value[i], dp[i-1][j])
			}
		}
	}

	return dp[len(weight)-1][cap]
}

// 状态压缩：根据递推公式 dp[i][j] = max(dp[i-1][j-weight[i]] + value[i], dp[i-1][j])可以知道，当前的值之和上一行的值有关，准确来说
// 仅仅和左上角的值有关，那么其实这个状态可以压缩到一行。即dp[j] = max(dp[j-weight[i]] + value[i], dp[j])
// 状态转移方程： dp[j] = max(dp[j-weight[i]] + value[i], dp[j])
// 初始化：使用物品0初始化第一行
// 遍历顺序：从前往后，只能先遍历物品，在遍历背包
// dp大小： dp := make([]int, cpa+1)
// 返回值： dp[cap]
func BaggageMaxValueDp03(weight, value []int, cap int) int {
	dp := make([]int, cap+1)

	for j := weight[0]; j <= cap; j++ {
		dp[j] = value[0]
	}

	fmt.Println(dp)
	for i := 1; i < len(weight); i++ {
		// TODO 这里必须倒序遍历，否则一个物品会被使用多次
		for j := cap; j >= weight[i]; j-- {
			dp[j] = max(dp[j-weight[i]]+value[i], dp[j])
		}
		fmt.Println(dp)
	}

	return dp[cap]
}

func TestGetMaxValue(t *testing.T) {
	var testData = []struct {
		weight []int
		value  []int
		cap    int
		want   int
	}{
		{weight: []int{1, 3, 4}, value: []int{15, 20, 30}, cap: 4, want: 35},
	}

	for _, tt := range testData {
		get := BaggageMaxValueDp03(tt.weight, tt.value, tt.cap)
		if get != tt.want {
			t.Fatalf("weight:%v, value:%v, cap:%v, want:%v, get:%v", tt.weight, tt.value, tt.cap, tt.want, get)
		}
	}
}
