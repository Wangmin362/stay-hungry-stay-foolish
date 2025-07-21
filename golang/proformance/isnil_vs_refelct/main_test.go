package isnil_vs_refelct

import (
	"reflect"
	"testing"
)

type Person struct {
	age     int
	name    string
	friends []string
	sex     string
}

func isNil(node *Person) bool {
	return node == nil
}

func reflectNil(node *Person) bool {
	return reflect.ValueOf(node).IsNil()
}

// 基准测试 isNil 方法
func BenchmarkIsNil(b *testing.B) {
	// 创建一个 Person 实例
	var person *Person
	// 使用 b.N 来执行多次基准测试
	for i := 0; i < b.N; i++ {
		// 调用 isNil 方法
		isNil(person)
	}
}

// 基准测试 reflectNil 方法
func BenchmarkReflectNil(b *testing.B) {
	// 创建一个 Person 实例
	var person *Person
	// 使用 b.N 来执行多次基准测试
	for i := 0; i < b.N; i++ {
		// 调用 reflectNil 方法
		reflectNil(person)
	}
}
