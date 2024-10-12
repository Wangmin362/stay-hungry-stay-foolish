package _0_basic

import "testing"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	res, cur, nxt := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nxt = max(nxt, i+nums[i])
		if i == cur { // 走到了当前覆盖范围的尽头，需要走下一个范围
			res++ // 所以需要跳一次
			cur = nxt
			if cur >= len(nums)-1 { // 说明可以走到尽头
				break
			}
		}
	}
	return res
}

func TestJump(t *testing.T) {

}
