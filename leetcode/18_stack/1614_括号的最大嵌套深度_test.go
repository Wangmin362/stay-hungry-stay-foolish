package _1_array

func maxDepth(s string) int {
	stack := make([]byte, 0, len(s))

	var res int
	idx := 0
	for ; idx < len(s); idx++ {
		if s[idx] == '(' {
			stack = append(stack, '(')
			res = max(res, len(stack))
		} else if s[idx] == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
