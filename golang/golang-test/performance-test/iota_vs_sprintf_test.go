package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// 57.33 ns/op
func Benchmark_Sprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

// 26.06 ns/op
func Benchmark_Itoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(int(rand.Int31()))
	}
}

// 26.08 ns/op
func Benchmark_FormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(rand.Int31()), 10)
	}
}
