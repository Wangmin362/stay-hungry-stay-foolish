package _0_basic

func rangeBitwiseAnd(left int, right int) int {
	// 这个循环的目的是为了找到 left 和 right 在二进制表示上相同的前缀部分。按位与操作会在高位上保留相同的值，
	// 而在低位上，若有任何一个数字为 0，则结果为 0。
	for right > left {
		// 这个操作会将 right 的最低位的 1 变成 0。这样做的效果是逐渐缩小 right，直到它与 left 共享相同的前缀。
		// 例如，如果 right 是 111（二进制）并且 left 是 101，第一次迭代将 right 改为 110，然后 101 和 110 在下一次迭代将变成 100，直到它们的高位部分相同。
		right &= right - 1
	}
	return left & right
}
