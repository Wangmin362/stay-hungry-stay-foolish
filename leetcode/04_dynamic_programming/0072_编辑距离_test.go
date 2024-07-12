package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/edit-distance/description/

func minDistanceII(word1 string, word2 string) int {
	// dp[i][j]定义为word1[:i-1]转换为word2[:j-1]最少的修改次数
	// word1[i-1]  = word2[j-1] dp[i][j] = dp[i-1][j-1]
	// word1[i-1] != word2[j-1] word1转为word2可以有删除、插入、以及替换操作
	//    删除/插入：dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1]+1, dp[i-1][j-1]+2)
	//    替换   dp[i][j] = dp[i-1][j-1] + 1
	dp := make([][]int, len(word1)+1)
	// dp[0][j] = j
	// dp[i][0] = i
	// dp[0][0] = 0
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func TestMinDistanceII(t *testing.T) {
	//fmt.Println(minDistanceII("horse", "ros"))
	fmt.Println(minDistanceII("intention", "execution"))
}
