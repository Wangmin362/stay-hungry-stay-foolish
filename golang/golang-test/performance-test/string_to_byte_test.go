package main

import (
	"io/ioutil"
	"testing"
)

// 结论：尽量减少字符串转为byte数字，性能损耗相当大

// 12.47 ns/op
func Benchmark_String2Bytes(b *testing.B) {
	data := "Hello world"
	w := ioutil.Discard
	for i := 0; i < b.N; i++ {
		w.Write([]byte(data))
	}
}

// 1.016 ns/op
func Benchmark_Bytes(b *testing.B) {
	data := []byte("Hello world")
	w := ioutil.Discard
	for i := 0; i < b.N; i++ {
		w.Write(data)
	}
}
