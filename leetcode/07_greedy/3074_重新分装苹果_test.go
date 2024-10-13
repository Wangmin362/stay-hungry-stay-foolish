package _0_basic

import (
	"sort"
	"testing"
)

func minimumBoxes(apple []int, capacity []int) int {
	var totalApple int
	for _, app := range apple {
		totalApple += app
	}

	sort.Ints(capacity)
	var res int
	for i := len(capacity) - 1; i >= 0; i-- {
		res++
		totalApple -= capacity[i]
		if totalApple <= 0 {
			break
		}
	}

	return res
}

func TestMinimumBoxes(t *testing.T) {
	var testsdata = []struct {
		apple    []int
		capacity []int
		want     int
	}{
		{apple: []int{1, 3, 2}, capacity: []int{4, 3, 1, 5, 2}, want: 2},
	}
	for _, tt := range testsdata {
		get := minimumBoxes(tt.apple, tt.capacity)
		if get != tt.want {
			t.Fatalf("apple:%v, capacity:%v, want:%v, get:%v", tt.apple, tt.capacity, tt.want, get)
		}
	}
}
