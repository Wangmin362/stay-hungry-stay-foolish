package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// 先来看看使用泛型的时候，对于每一种类型，我们都需要定义一个函数支持
func SumInt(x, y int) int {
	return x + y
}
func TestSumInt(t *testing.T) {
	assert.Equal(t, SumInt(3, 5), 8, "not equal")
}
func SumFloat(x, y float32) float32 {
	return x + y
}
func TestSumFloat(t *testing.T) {
	assert.Equal(t, SumFloat(3.2, 5.7), 8.9, "not equal")
}

// 如果使用泛型

func Sum[T int | int64 | float32 | string](x, y T) T {
	return x + y
}
func TestSum(t *testing.T) {
	assert.Equal(t, Sum[int](3, 5), 8, "int not equal")
	assert.Equal(t, Sum[int64](int64(3), int64(5)), int64(8), "int64 not equal")
	assert.Equal(t, Sum[float32](float32(3.2), float32(5.7)), float32(8.9), "float32 not equal")
	assert.Equal(t, Sum[string]("hello", "world"), "helloworld", "string not equal")
}
