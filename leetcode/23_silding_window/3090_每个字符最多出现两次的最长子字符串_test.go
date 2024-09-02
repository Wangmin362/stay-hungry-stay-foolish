package _1_array

// https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/description/

func maximumLengthSubstring(s string) int {
	cache := [26]int{}
	ans, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cache[c]++
		for cache[c] > 2 {
			cache[s[left]-'a']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
