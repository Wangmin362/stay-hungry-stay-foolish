package _0_basic

// 参考视频：https://www.bilibili.com/video/BV1Qg411q7ia/?vd_source=d039ae9ec8b71e411a906e821301b7ac

// 方法一，使用前后缀分贝保存前缀最大值，每个格子看成有一个水桶节水，分别取左边的最大高度和右边的最大高度，这两个取最小值
func trap01(height []int) int {
	if len(height) < 3 {
		return 0
	}
	preMax, sufMax := make([]int, len(height)), make([]int, len(height))
	// 计算前缀最大值
	preMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		preMax[i] = max(height[i], preMax[i-1])
	}

	// 计算后缀最大值
	sufMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		sufMax[i] = max(height[i], sufMax[i+1])
	}

	var res int
	for i := 0; i < len(height); i++ {
		high := min(preMax[i], sufMax[i])
		res += high - height[i] // 因为宽度都是一，所以高度就是面积
	}

	return res
}

// 双指针
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	res, left, right := 0, 0, len(height)-1
	preMax, sufMax := 0, 0
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			res += preMax - height[left]
			left += 1
		} else {
			res += sufMax - height[right]
			right -= 1
		}
	}

	return res
}
