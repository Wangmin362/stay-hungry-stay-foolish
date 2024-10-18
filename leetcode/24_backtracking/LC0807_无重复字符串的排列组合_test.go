package _1_array

import "testing"

func permutation(S string) []string {
	var backtracking func()

	var res []string
	var path []byte
	used := make([]bool, len(S))
	backtracking = func() {
		if len(path) == len(S) {
			res = append(res, string(path))
			return
		}

		for i := 0; i < len(S); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, S[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtracking()
	return res
}

func TestPermutation(t *testing.T) {
	t.Log(permutation("abc"))
}
