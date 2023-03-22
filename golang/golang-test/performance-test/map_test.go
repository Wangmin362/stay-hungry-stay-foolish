package main

import "testing"

// 结论：对于map也是同理，能指定容量的情况下，尽可能的指定初始化容量，避免后续经常扩容，导致性能下降

var mapSize = 1000

// 38684 ns/op
func Benchmark_MapNoCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make(map[int]int)
		for k := 0; k < mapSize; k++ {
			data[k] = k
		}
	}
}

// 17143 ns/op
func Benchmark_MapCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make(map[int]int, mapSize)
		for k := 0; k < mapSize; k++ {
			data[k] = k
		}
	}
}
