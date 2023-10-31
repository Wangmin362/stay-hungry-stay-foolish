package main

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int

func Test03(t *testing.T) {
	var i int
	var j MyInt

	i = int(j) // 必须强转

	ti := reflect.TypeOf(i)
	fmt.Println("type of i:", ti.String())

	tj := reflect.TypeOf(j)
	fmt.Println("type of j:", tj.String())

	fmt.Println("kind of i:", ti.Kind())
	fmt.Println("kind of j:", tj.Kind())
}
