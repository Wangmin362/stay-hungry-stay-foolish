package main

import (
	"testing"
)

func TestUploadLog(t *testing.T) {
	var testdata = []struct {
		logs []int
		want int
	}{
		{logs: []int{1, 2, 92, 3}, want: 91},
		{logs: []int{3, 7, 40, 10, 60}, want: 37},
	}
	for _, tt := range testdata {
		get := firstUploadLog(tt.logs)
		if get != tt.want {
			t.Fatalf("logs:%v, want:%v, get:%v", tt.logs, tt.want, get)
		}
	}
}
