package _0_basic

import (
	"fmt"
	"testing"
)

func canCompleteCircuit(gas []int, cost []int) int {
	total, sum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		sum += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}

	return start
}

func TestCanCompleteCircuit(t *testing.T) {
	//fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
	fmt.Println(canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))
}
