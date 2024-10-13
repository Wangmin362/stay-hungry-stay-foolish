package _0_basic

import (
	"reflect"
	"sort"
	"testing"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		if boxTypes[i][1] == boxTypes[j][1] {
			return boxTypes[i][0] > boxTypes[j][0]
		}
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var res int
	for i := 0; i < len(boxTypes) && truckSize > 0; {
		if boxTypes[i][0] <= 0 {
			i++
			continue
		}
		res += boxTypes[i][1]
		boxTypes[i][0]--
		truckSize--
	}

	return res
}

func TestMaximumUnits(t *testing.T) {
	var testsdata = []struct {
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{boxTypes: [][]int{{1, 3}, {2, 2}, {3, 1}}, truckSize: 4, want: 8},
	}
	for _, tt := range testsdata {
		get := maximumUnits(tt.boxTypes, tt.truckSize)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("boxTypes:%v, truckSize:%v, want:%v, get:%v", tt.boxTypes, tt.truckSize, tt.want, get)
		}
	}
}
