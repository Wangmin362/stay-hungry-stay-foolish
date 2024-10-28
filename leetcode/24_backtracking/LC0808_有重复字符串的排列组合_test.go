package _1_array

import "sort"

func permutation0807(S string) []string {
	bytes := []byte(S)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	var backtracking func()
	var res []string
	used := make([]bool, len(bytes))
	var path []byte
	backtracking = func() {
		if len(path) == len(S) {
			res = append(res, string(path))
			return
		}
		rowUsed := make(map[byte]struct{})
		for i := 0; i < len(bytes); i++ {
			if _, ok := rowUsed[bytes[i]]; used[i] || ok {
				continue
			}

			rowUsed[bytes[i]] = struct{}{}
			used[i] = true
			path = append(path, bytes[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking()
	return res
}
