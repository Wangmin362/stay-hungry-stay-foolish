package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/target-sum/description/

// 也可以抽象为01背包问题，每个元素只能取一次。把数组分为两个组，分别是left, right。根据题目的意思left - right = 3 并且left + right = sum
// 那么left - (sum -left) = 3，即2left - sum = 3，即left = (sum + 3) / 2，因此left = (sum + target) /2
func findTargetSumWays(nums []int, target int) int {
	// dp[j]定义为容量为j的背包，有多少种方法可以装满，
	// 根据上面的定义，dp[j] = dp[j-1]+dp[j-2]+...dp[0]
	// dp[j] += dp[j - nums[i]]
	sum := 0
	for idx := range nums {
		sum += nums[idx]
	}
	if target > sum {
		return 0
	}
	if target < 0 && -target > sum {
		return 0
	}

	if (sum+target)%2 == 1 { // 说明无法拼凑
		return 0
	}

	left := (sum + target) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
		fmt.Println(dp)
	}

	return dp[left]
}

// 题目分析: 其实就是把数组分为两堆，左边一堆取正号，右边一堆取负号。 两堆数字总和的差集就是target，因此有
// left + right = sum   left - right = target  ==> left + (left - target) = sum  => left = (sum + target) / 2
// 这个left其实就是我们再数组中找到能够凑成left总和的数组有多少种方法。
// 问题抽象： 再Nums书中中找到总和为left有多少种不同的方法
// 明确定义： dp[j]表示从nums数组的0..i个数组中有多少种凑成容量为j的方法
// 状态转移公式： 可以通过举例的方式明确公式，当j=4的时候
// nums[i]=0   dp[4]
// nums[i]=1   dp[3]
// nums[i]=2   dp[2]
// nums[i]=3   dp[1]
// nums[i]=4   dp[0]
// 因此 dp[j] += dp[j-nums[i]]
// 初始化dp[j] = 1
// 遍历顺序：先便利物品，再遍历容量，容量倒序遍历

func findTargetSumWays0912(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if target > sum {
		return 0
	}
	if target < 0 && -target > sum {
		return 0
	}

	if (sum+target)%2 == 1 { // 说明凑不出来这样的target
		return 0
	}

	capacity := (sum + target) >> 1
	dp := make([]int, capacity+1)
	dp[0] = 1
	//fmt.Println(dp)
	for i := 0; i < len(nums); i++ {
		for j := capacity; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
		//fmt.Println(dp)
	}

	return dp[capacity]
}

// 本地最终可以转换为01背包问题
// 若求的是最大价值，那么dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-weight[i])+value[i])
// 若求的是最小价值，那么dfs(i, c) = min(dfs(i-1, c), dfs(i-1, c-weight[i])+value[i])
// 若求的数量，显然是二者之和：dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-weight[i])
// 可以想到 dfs(i, c)表示从前i个物品选择一些物品装满容量为c的背包的最大数量，自然而然就等于最后一个物品选和不选的最大数量之和
// 也就是说本题的递归方程为：dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-weight[i])
// left + right = sum;  left-right=target ==> left + (left - target) = sum ==> left = (sum + target) / 2
// 也就是说容量为：capacity = (sum + target) / 2
func findTargetSumWaysDfs(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}

	capacity := (sum + target) >> 1
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 { // 如果容量为0，说明恰好装满了，此时自然需要返回1
				return 1
			}
			return 0
		}
		if c < nums[i] { // 若剩余的容量放不下当前物品，只能选择不选这个物品
			res := dfs(i-1, c)
			return res
		}

		res := dfs(i-1, c) + dfs(i-1, c-nums[i])
		return res
	}

	return dfs(len(nums)-1, capacity)
}

// 修改为记忆化搜索
func findTargetSumWaysDfsMem(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}

	capacity := (sum + target) >> 1

	mem := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mem[i] = make([]int, capacity+1)
	}
	// 全部初始化为-1，因为结果不可能是-1
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= capacity; j++ {
			mem[i][j] = -1
		}
	}

	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 { // 如果容量为0，说明恰好装满了，此时自然需要返回1
				return 1
			}
			return 0
		}

		if mem[i][c] != -1 {
			return mem[i][c]
		}

		if c < nums[i] { // 若剩余的容量放不下当前物品，只能选择不选这个物品
			res := dfs(i-1, c)
			mem[i][c] = res
			return res
		}

		res := dfs(i-1, c) + dfs(i-1, c-nums[i])
		mem[i][c] = res
		return res
	}

	return dfs(len(nums)-1, capacity)
}

// 修改为地推，也就是动态规划
// 原始公式为：dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-weight[i])
// 修改为递推：f[i][c] = f[i-1][c] + f[i-1][c-weight[i]]
// 由于索引有负数，因此两边同时加一，于是递推公式为：f[i+1][c] = f[i][c] + f[i][c-weight[i]]
func findTargetSumWaysDP(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}

	capacity := (sum + target) >> 1

	f := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		f[i] = make([]int, capacity+1)
	}
	f[0][0] = 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= capacity; j++ {
			if j < nums[i] {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = f[i][j] + f[i][j-nums[i]]
			}
		}
	}
	return f[len(nums)][capacity]
}

// 修改为地推，也就是动态规划
// 原始公式为：dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-weight[i])
// 修改为递推：f[i][c] = f[i-1][c] + f[i-1][c-weight[i]]
// 由于索引有负数，因此两边同时加一，于是递推公式为：f[i+1][c] = f[i][c] + f[i][c-weight[i]]
// 空间优化，优化到两个数组
func findTargetSumWaysDPOpt(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}

	capacity := (sum + target) >> 1

	f := make([][]int, 2)
	f[0], f[1] = make([]int, capacity+1), make([]int, capacity+1)
	f[0][0] = 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= capacity; j++ {
			if j < nums[i] {
				f[(i+1)%2][j] = f[i%2][j]
			} else {
				f[(i+1)%2][j] = f[i%2][j] + f[i%2][j-nums[i]]
			}
		}
	}
	return f[len(nums)%2][capacity]
}

// 修改为地推，也就是动态规划
// 原始公式为：dfs(i, c) = dfs(i-1, c) + dfs(i-1, c-weight[i])
// 修改为递推：f[i][c] = f[i-1][c] + f[i-1][c-weight[i]]
// 由于索引有负数，因此两边同时加一，于是递推公式为：f[i+1][c] = f[i][c] + f[i][c-weight[i]]
// 再次空间优化，优化到一个数组 f[i+1][c] = f[i][c] + f[i][c-weight[i]] ==> f[c] = f[c] + f[c-weight[i]]
func findTargetSumWaysDPOpt2(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}

	capacity := (sum + target) >> 1

	f := make([]int, capacity+1)
	f[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := capacity; j >= 0; j-- {
			if j < nums[i] {
				f[j] = f[j]
			} else {
				f[j] = f[j] + f[j-nums[i]]
			}
		}
	}
	return f[capacity]
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays0912([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWaysDP([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWaysDPOpt([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWaysDPOpt2([]int{1, 1, 1, 1, 1}, 3))
}
