package _1_array

import (
	"fmt"
	"testing"
)

func lastRemaining(n int) int {
	arr := make([]bool, n)
	var del int
	for del < n-1 {
		cnt := 0
		for i := 0; i < n; i++ {
			if arr[i] {
				continue
			}
			cnt++
			if cnt == 2 {
				arr[i] = true
				cnt = 0
				del++
			}
		}
	}
	for i := 0; i < n; i++ {
		if arr[i] {
			return i + 1
		}
	}
	return -1
}

func TestLastRemaining(t *testing.T) {
	fmt.Println(lastRemaining(9))
}
