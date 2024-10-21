package _0_basic

// 思路：每次比较右上角的值，要么去除一行，要么去除一列
func searchMatrix(matrix [][]int, target int) bool {
	top, right := 0, len(matrix[0])-1 // 右上角
	for top < len(matrix) && right >= 0 {
		val := matrix[top][right]
		if val == target {
			return true
		} else if val > target { // 当前列可以去除掉
			right--
		} else { // 这一行可以去除掉
			top++
		}
	}
	return false
}
