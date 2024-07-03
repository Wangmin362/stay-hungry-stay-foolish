package _1_array

// https://leetcode.cn/problems/combination-sum-iii/description/

func combinationSum3(k int, n int) [][]int {
	var backtracking func(k, n, startIdx, sum int)

	var res [][]int
	var path []int
	backtracking = func(k, n, startIdx, sum int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := startIdx; i <= 9; i++ {
			if 9-i+1 < k-len(path) || sum > n { // 回溯剪枝
				break
			}
			path = append(path, i)
			backtracking(k, n, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtracking(k, n, 1, 0)
	return res
}
