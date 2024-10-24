package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/last-stone-weight-ii/description/

// 此题可以抽象为01背包问题，题目的核心其实是要把石头分成两堆，这两对石头的重要要尽可能的相等，要整两队石头相撞之后，剩下的就是最小可能的重量。
// 因此可以把题目转换为求容量为sum/2的背包最多可以装的石头的重量（价值），此题当中，石头的重量为石头的价值。左边石头堆的重量为dp[sum/2]，
// 右边石头堆的重量为sum - dp[sum/2]，因此最小可能的重量为sum - dp[sum/2] - dp[sum/2] = sum - 2*dp[sum/2]
func lastStoneWeightII(stones []int) int {
	// dp[j]定义为容量为j的背包最多可以装的石头的价值，其实价值就是重量
	// dp[j] = max(dp[j], dp[j-weight[i]] + value[i]), 在本题当中，dp[j] = max(dp[j], dp[j - stones[i]] + stones[i])
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}

	sum := 0
	for idx := range stones {
		sum += stones[idx]
	}

	leftWeight := sum >> 1

	dp := make([]int, leftWeight+1)
	// 初始化第一行，即如果只有stones[0]的情况下，这些不同容量的背包最多可以装的石头的最大重量，只要背包的容量超过石头的重量，最多其实就是可以装
	// stones[0]的重量，因此此时就只有一个石头
	for j := stones[0]; j <= leftWeight; j++ {
		dp[j] = stones[0]
	}
	fmt.Println(dp)
	for i := 1; i < len(stones); i++ {
		for j := leftWeight; j > 0; j-- {
			if j < stones[i] { // 当前背包的容量本来就比石头的重量小，那么肯定放不进去，因此只能丢弃
				dp[j] = dp[j]
			} else {
				noput := dp[j]
				put := dp[j-stones[i]] + stones[i]
				dp[j] = int(math.Max(float64(noput), float64(put)))
			}
		}
		//fmt.Println(dp)
	}

	left := dp[leftWeight]
	right := sum - left
	return right - left
}

func lastStoneWeightII02(stones []int) int {
	// dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
	sum := 0
	for idx := range stones {
		sum += stones[idx]
	}
	mid := sum >> 1
	dp := make([]int, mid+1)
	for i := 0; i < len(stones); i++ {
		for j := mid; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
		fmt.Println(dp)
	}

	left := dp[mid]
	right := sum - left
	return right - left
}

// 计算总和，遍历所有的可能，计算差值，最小的差值就是结果
func lastStoneWeightIIBacktracking(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	var backtracking func(start, cnt int)

	res := sum
	backtracking = func(start, cnt int) {
		if cnt > sum-cnt {
			res = min(res, cnt-(sum-cnt))
		} else {
			res = min(res, sum-cnt-cnt)
		}

		for i := start; i < len(stones); i++ {
			backtracking(i+1, cnt+stones[i])
		}
	}

	backtracking(0, 0)
	return res
}

// 题目分析：其实就是尽可能找到两堆相等的石头，结果就是这两堆尽可能共享的石头的差值
// 抽象模型：每个石头只能取一次，我只需要使用容量为sum/2的背包，然后使用这些石头装，找到背包的最大价值，然后两边相减即可，这样就抽象出来了
// 01背包问题
// 明确定义：dp[j]为前i个石头，也即是0..i个石头可以转入背包容量为j的最大重量，石头的重量为stones[i], 价值也是stones[i]
// 状态方程：dp[j] = max(dp[j-stones[i]] + stones[i], dp[j])
// 初始化：使用第一块石头初始化
// 遍历顺序：先物品，再容量，容量需要倒叙遍历，防止物品取了多次
// dp数组大小： sum/2+1
// 返回值：(sum - sum/2) - dp[sum/2]
func lastStoneWeightII0912(stones []int) int {
	sum := 0
	for _, sto := range stones {
		sum += sto
	}

	capacity := sum >> 1
	dp := make([]int, capacity+1)
	for j := stones[0]; j <= capacity; j++ {
		dp[j] = stones[0]
	}

	for i := 1; i < len(stones); i++ {
		for j := capacity; j >= stones[i]; j-- {
			dp[j] = max(dp[j-stones[i]]+stones[i], dp[j])
		}
	}

	return (sum - dp[capacity]) - dp[capacity]
}

// 题目分析：01背包问题，只要求凑出总和一半的背包就行
// 递归：dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-stones[i])+stones[i])
func lastStoneWeightIIDfs(stones []int) int {
	var sum int
	for _, num := range stones {
		sum += num
	}
	capacity := sum >> 1
	var dfs func(i, c int) int
	mem := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		mem[i] = make([]int, capacity+1)
		for j := 0; j <= capacity; j++ {
			mem[i][j] = -1
		}
	}

	dfs = func(i, c int) int {
		if i < 0 {
			return 0
		}
		if mem[i][c] != -1 {
			return mem[i][c]
		}

		if c < stones[i] { // 如果背包的剩余容量小于石头的重量，那么这块石头一定放不进去
			res := dfs(i-1, c)
			mem[i][c] = res
			return res
		}
		res := max(dfs(i-1, c), dfs(i-1, c-stones[i])+stones[i])
		mem[i][c] = res
		return res
	}

	half := dfs(len(stones)-1, capacity)
	other := sum - half
	return other - half
}

// 题目分析：01背包问题，只要求凑出总和一半的背包就行
// 递归：dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-stones[i])+stones[i])
// 递推：f[i][c] = max(f[i-1][c], f[i-1][c-stones[i]]+stones[i])
// 两边同时加一，可以得到：
// 递推：f[i+1][c] = max(f[i][c], f[i][c-stones[i]]+stones[i])
func lastStoneWeightIIDp(stones []int) int {
	var sum int
	for _, num := range stones {
		sum += num
	}
	capacity := sum >> 1

	f := make([][]int, len(stones)+1)
	for i := 0; i <= len(stones); i++ {
		f[i] = make([]int, capacity+1)
	}

	for i := 0; i < len(stones); i++ {
		for j := 0; j <= capacity; j++ {
			if j < stones[i] {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j], f[i][j-stones[i]]+stones[i])
			}
		}
	}
	half := f[len(stones)][capacity]
	other := sum - half
	return other - half
}

// 题目分析：01背包问题，只要求凑出总和一半的背包就行
// 递归：dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-stones[i])+stones[i])
// 递推：f[i][c] = max(f[i-1][c], f[i-1][c-stones[i]]+stones[i])
// 两边同时加一，可以得到：
// 递推：f[i+1][c] = max(f[i][c], f[i][c-stones[i]]+stones[i])
// 优化空间到O(2*len(stones))
// 递推：f[(i+1)%2][c] = max(f[i%2][c], f[i%2][c-stones[i]]+stones[i])
func lastStoneWeightIIDpOp1(stones []int) int {
	var sum int
	for _, num := range stones {
		sum += num
	}
	capacity := sum >> 1

	f := make([][]int, 2)
	f[0] = make([]int, capacity+1)
	f[1] = make([]int, capacity+1)

	for i := 0; i < len(stones); i++ {
		for j := 0; j <= capacity; j++ {
			if j < stones[i] {
				f[(i+1)%2][j] = f[i%2][j]
			} else {
				f[(i+1)%2][j] = max(f[i%2][j], f[i%2][j-stones[i]]+stones[i])
			}
		}
	}
	half := f[len(stones)%2][capacity]
	other := sum - half
	return other - half
}

// 题目分析：01背包问题，只要求凑出总和一半的背包就行
// 递归：dfs(i, c) = max(dfs(i-1, c), dfs(i-1, c-stones[i])+stones[i])
// 递推：f[i][c] = max(f[i-1][c], f[i-1][c-stones[i]]+stones[i])
// 两边同时加一，可以得到：
// 递推：f[i+1][c] = max(f[i][c], f[i][c-stones[i]]+stones[i])
// 优化空间到O(2*len(stones))
// 递推：f[(i+1)%2][c] = max(f[i%2][c], f[i%2][c-stones[i]]+stones[i])
// 继续优化空间到O(len(stones))
// 递推：f[c] = max(f[c], f[c-stones[i]]+stones[i])
func lastStoneWeightIIDpOp2(stones []int) int {
	var sum int
	for _, num := range stones {
		sum += num
	}
	capacity := sum >> 1

	f := make([]int, capacity+1)

	for i := 0; i < len(stones); i++ {
		for j := capacity; j >= stones[i]; j-- {
			f[j] = max(f[j], f[j-stones[i]]+stones[i])
		}
	}
	half := f[capacity]
	other := sum - half
	return other - half
}

func TestLastStoneWeightII(t *testing.T) {
	var testdata = []struct {
		stones []int
		want   int
	}{
		//{stones: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{stones: []int{31, 26, 33, 21, 40}, want: 5},
	}
	for _, tt := range testdata {
		get := lastStoneWeightIIDpOp2(tt.stones)
		if get != tt.want {
			t.Fatalf("stones:%v, want:%v, get:%v", tt.stones, tt.want, get)
		}
	}
}
