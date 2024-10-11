package main

import (
	"reflect"
	"testing"
)

func TestEscapeEarth(t *testing.T) {
	testdata := []struct {
		total       int
		starts      []pair
		wantCnt     int
		wantEngines []int
	}{
		//{total: 8, starts: []pair{{0, 2}, {0, 6}}, wantCnt: 2, wantEngines: []int{0, 4}},
		{total: 8, starts: []pair{{0, 0}, {1, 7}}, wantCnt: 1, wantEngines: []int{4}},
	}

	for _, tt := range testdata {
		getCnt, getEngines := escapeEarth02(tt.total, tt.starts)
		if getCnt != tt.wantCnt || !reflect.DeepEqual(getEngines, tt.wantEngines) {
			t.Fatalf("wantCnt:%v, wantEngines:%v, getCnt:%v, getEngines:%v", tt.wantCnt, tt.wantEngines, getCnt, getEngines)
		}
	}
}
