package _0_basic

// https://leetcode.cn/problems/powx-n/?envType=study-plan-v2&envId=top-interview-150

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		// 如果 n 是负数，将 x 变为 1/x，并将 n 取反，以处理负次幂。
		x = 1 / x
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n%2 == 1 { // 如果 n 是奇数
			result *= x
		}
		x *= x // 平方 x
		n /= 2 // 除以 2
	}
	return result
}
