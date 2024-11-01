package _1_array

func gameResult(head *ListNode) string {
	even, odd := head, head.Next
	evenScore, oddScore := 0, 0
	for even != nil && odd != nil {
		if even.Val > odd.Val {
			evenScore++
		} else if even.Val < odd.Val {
			oddScore++
		}
		even = odd.Next
		if even == nil {
			break
		}
		odd = even.Next
	}

	if evenScore > oddScore {
		return "Even"
	} else if evenScore < oddScore {
		return "Odd"
	}

	return "Tie"
}
