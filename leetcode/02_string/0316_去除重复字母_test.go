package _1_array

import "sort"

func removeDuplicateLetters(s string) string {
	cache := [26]int{}
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
		if cache[s[i]-'a'] == 1 {
			res = append(res, s[i])
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return string(res)
}
