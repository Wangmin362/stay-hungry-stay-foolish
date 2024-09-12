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
	fmt.Println(dp)
	for i := 0; i < len(nums); i++ {
		for j := capacity; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
		fmt.Println(dp)
	}

	return dp[capacity]
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays0912([]int{100}, -200))
}
