package type_assert

import "testing"

import (
	"fmt"
)

func TestGenericTypeAssert(t *testing.T) {
	var i interface{} = "hello"

	s, ok := i.(string) // 带有两个返回值的类型断言

	if ok {
		fmt.Println(s) // 输出: hello
	} else {
		fmt.Println("Not a string")
	}

	// 使用另一种类型断言
	f, ok := i.(float64)
	if ok {
		fmt.Println(f)
	} else {
		fmt.Println("Not a float64")
	}
}
