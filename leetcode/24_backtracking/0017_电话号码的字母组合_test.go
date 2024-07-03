package _1_array

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

// 分析：递归深度，其实就是字符串的长度
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	digMap := map[uint8][]uint8{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var backtracking func(digits string, startIdx int)

	var res []string
	var path string
	backtracking = func(digits string, startIdx int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}
		for _, c := range digMap[digits[startIdx]] { // 这里无法剪枝，因为每个分支都是符合要求的
			path += string(c)
			backtracking(digits, startIdx+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(digits, 0)
	return res
}
