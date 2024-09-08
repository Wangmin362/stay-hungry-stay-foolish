package _0_basic

// https://leetcode.cn/problems/find-the-difference/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func findTheDifference(s string, t string) byte {
	var res byte
	for _, c := range s {
		res ^= byte(c)
	}

	for _, c := range t {
		res ^= byte(c)
	}

	return res
}
