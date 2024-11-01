package _1_array

type node7 struct {
	Val  int
	Next *node7
	Prev *node7
}

func toArray(head *node7) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
