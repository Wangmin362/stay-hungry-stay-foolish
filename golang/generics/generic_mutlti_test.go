package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// TODO 1. 如果由多个不同类型的泛型参数如何表达？

func SumMulti[A int | float32, B int | float32, C float64](a A, b B) C {
	return C(a + A(b))
}

func TestMulit(t *testing.T) {
	assert.Equal(t, SumMulti[int, float32, float64](5, float32(3.2)), 8, "not equal")
}
