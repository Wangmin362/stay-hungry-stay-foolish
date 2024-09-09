package _0_basic

// https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

// 参考灵茶山艾府，其实就是统计不同元素的个数
func minimumOperations(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums {
		if num != 0 {
			m[num] = struct{}{}
		}
	}
	return len(m)
}
