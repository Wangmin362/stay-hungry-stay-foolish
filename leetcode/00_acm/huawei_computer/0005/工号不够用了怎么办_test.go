package main

import "testing"

func TestGongHao(t *testing.T) {
	var testdata = []struct {
		total int
		y     int
		want  int
	}{
		{total: 260, y: 1, want: 1},
		{total: 26, y: 1, want: 1},
		{total: 2600, y: 1, want: 2},
	}

	for _, tt := range testdata {
		get := gongHao(tt.total, tt.y)
		if get != tt.want {
			t.Fatalf("total:%v, y:%v, want:%v, get:%v", tt.total, tt.y, tt.want, get)
		}
	}
}
