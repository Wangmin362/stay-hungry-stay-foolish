package _1_array

import (
	"fmt"
	"testing"
)

func iceBreakingGame(num int, target int) int {
	nums := make([]bool, num)
	idx := 0
	del := 0
	cnt := 0
	for del < num-1 {
		if nums[idx%len(nums)] {
			idx++
			continue
		}
		cnt++
		if cnt == target {
			nums[idx%len(nums)] = true
			cnt = 0
			del++
		}
		idx++
	}

	for i := 0; i < num; i++ {
		if !nums[i] {
			return i
		}
	}
	return -1
}

func TestIceBreakingGame(t *testing.T) {
	fmt.Println(iceBreakingGame(7, 4))
}
