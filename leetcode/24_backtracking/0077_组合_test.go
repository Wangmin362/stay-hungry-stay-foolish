package _1_array

// https://leetcode.cn/problems/combinations/description/

func combine(n int, k int) [][]int {
	var backtracking func(n, k, startIdx int)

	var res [][]int // 存放最终结果
	var path []int  // 存放当前路径
	backtracking = func(n, k, startIdx int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			for idx, item := range path {
				tmp[idx] = item
			}
			res = append(res, tmp)
			return
		}
		for i := startIdx; i <= n; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)
	return res
}

// 回溯剪枝
func combine01(n int, k int) [][]int {
	var backtracking func(n, k, startIdx int)

	var res [][]int // 存放最终结果
	var path []int  // 存放当前路径
	backtracking = func(n, k, startIdx int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			for idx, item := range path {
				tmp[idx] = item
			}
			res = append(res, tmp)
			return
		}
		//k - len(path) 还剩k - len(path)个元素要取
		// n - (k - len(path)) 计算索引范围
		//	譬如：n=4,k=3 path=0  n - (k - len(path)) = 1 所以要加1
		for i := startIdx; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)
	return res
}

// 回溯剪枝  优化拷贝和剪枝理解
func combine03(n int, k int) [][]int {
	var backtracking func(n, k, startIdx int)

	var res [][]int // 存放最终结果
	var path []int  // 存放当前路径
	backtracking = func(n, k, startIdx int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIdx; i <= n; i++ {
			// n - i + 1 当前for循环还可以搜索递归层数
			// k - len(path) 当前还需要的数字
			// 如果需要的递归层数大于还需要的数组，那么是合法的，如果小于的化肯定无法满足条件，直接舍弃
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)
	return res
}
