package _0_basic

import (
	"sort"
	"testing"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	diff := make([]int, len(capacity))
	for i := 0; i < len(capacity); i++ {
		diff[i] = capacity[i] - rocks[i] // 还需要几块石头装满
	}

	sort.Ints(diff)
	res, idx := 0, 0
	for idx < len(diff) {
		if diff[idx] == 0 { // 说明当前背包已经是满的，不需要装
			res++
			idx++
		} else {
			break
		}
	}

	for additionalRocks > 0 && idx < len(diff) {
		if additionalRocks >= diff[idx] { // 说明当前剩余的石块数量大于装满背包需要的数量
			res++                        // 可以装满的背包加一
			additionalRocks -= diff[idx] // 维护剩余石块的装量
			idx++
		} else {
			break
		}
	}

	return res
}

func TestMaximumBags(t *testing.T) {
	var testsdata = []struct {
		capacity        []int
		rocks           []int
		additionalRocks int
		want            int
	}{
		{capacity: []int{2, 3, 4, 5}, rocks: []int{1, 2, 4, 4}, additionalRocks: 2, want: 3},
	}
	for _, tt := range testsdata {
		get := maximumBags(tt.capacity, tt.rocks, tt.additionalRocks)
		if get != tt.want {
			t.Fatalf("capacity:%v, rocks:%v, additionalRocks:%v, want:%v, get:%v",
				tt.capacity, tt.rocks, tt.additionalRocks, tt.want, get)
		}
	}
}
