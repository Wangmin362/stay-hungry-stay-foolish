package _1_array

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = -1 // 直接默认复制为-1
		}
	}

	l, r, t, b := 0, n-1, 0, m-1
	for l <= r && t <= b && head != nil {
		for i := l; i <= r && head != nil; i++ {
			res[t][i] = head.Val
			head = head.Next
		}
		t++
		if t > b || head == nil {
			break
		}

		for i := t; i <= b && head != nil; i++ {
			res[i][r] = head.Val
			head = head.Next
		}
		r--
		if r < l || head == nil {
			break
		}

		for i := r; i >= l && head != nil; i-- {
			res[b][i] = head.Val
			head = head.Next
		}
		b--
		if b < t || head == nil {
			break
		}

		for i := b; i >= t && head != nil; i-- {
			res[i][l] = head.Val
			head = head.Next
		}
		l++
		if l > r || head == nil {
			break
		}
	}

	return res
}
