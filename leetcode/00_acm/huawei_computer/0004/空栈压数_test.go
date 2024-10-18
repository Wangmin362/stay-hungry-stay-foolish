package main

import (
	"reflect"
	"testing"
)

func TestEmptyStack(t *testing.T) {
	var testdata = []struct {
		nums []int
		want []int
	}{
		{nums: []int{10, 20, 50, 80, 1, 1}, want: []int{2, 160}},
		{nums: []int{5, 10, 20, 50, 85, 1}, want: []int{1, 170}},
	}
	for _, tt := range testdata {
		get := emptyStack(tt.nums)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
