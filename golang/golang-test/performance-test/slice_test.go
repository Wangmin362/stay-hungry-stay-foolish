package main

import "testing"

// 结论：对于切片，应该尽量指定容量，否则底层数组经常扩容，势必对于性能有影响

var sliceSize = 1000

// 2563 ns/op
func Benchmark_SliceNoCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < sliceSize; k++ {
			data = append(data, k)
		}
	}
}

// 1002 ns/op
func Benchmark_SliceCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, sliceSize)
		for k := 0; k < sliceSize; k++ {
			data = append(data, k)
		}
	}
}
