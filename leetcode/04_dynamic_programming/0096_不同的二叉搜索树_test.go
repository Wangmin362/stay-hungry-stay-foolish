package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/unique-binary-search-trees/description/

func numTrees(n int) int {
	// dp[i] 为1到i个数字可以组成的二叉搜索树的数量
	// 1到i个数字可以组成的二叉搜索树可以分为头节点为1，头节点为2，头节点为3.....头节点为i的二叉搜索树
	// dp[3] = dp[0]*dp[2] + dp[1]*dp[1] + dp[2]*dp[0] 即[1,3]数组可以组成的搜查搜索的数量为头节点为1的二叉搜索树 + 头节点
	// 为2的二叉搜索树，以及头节点为3的二叉搜索树。头节点为1的二叉搜索树，左边的节点数量一定为0，右边的节点数量一定为2。 而头节点为2
	// 的二叉搜索树左边的节点数量一定为1，右边的节点数量一定为1。 头节点为3的二叉搜索树左边一定为2，右边的节点数量一定为0
	// dp[i] += dp[j] * dp[i-j-1] // 左边有j个节点， 右边一定有i-j-1个节点，因为头节点占用一个节点
	// dp[3] = dp[0]*dp[3-0-1] + dp[1]*dp[3-1-1] + dp[2]*dp[3-2-1]
	// dp[1] = dp[0]*dp[1-0-1] = dp[0]*dp[0]
	dp := make([]int, n+1)
	dp[0] = 1 // 空二叉树也是二叉搜索树
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ { // 头节点的左边节点可以是0个，也可以是1个节点，一直到i-1个节点
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

/*






sdfsdf




*/

func numTrees01(n int) int {
	// dp[4]其实就是头节点分别为1，2，3，4的情况之和、
	// 头节点为1  左边节点为0，右边节点为3个  因此dp[0] *dp[3]
	// 头节点为2  左边节点为1，右边节点为2个  因此dp[1] *dp[2]
	// 头节点为3  左边节点为2，右边节点为1个  因此dp[2] *dp[1]
	// 头节点为4  左边节点为3，右边节点为0个  因此dp[3] *dp[0]
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ { // 左边节点数量
			dp[i] += dp[j] * dp[i-j-1] // 头节点占用一个节点，因此还需要减一
		}
	}

	return dp[n]
}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees01(3))
}
