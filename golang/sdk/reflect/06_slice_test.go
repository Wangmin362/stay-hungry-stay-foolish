package main

import (
	"fmt"
	"reflect"
	"testing"
)

func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)

	fmt.Printf("%c", '[')
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Printf("%c\n", ']')
}

func TestSlice05(t *testing.T) {
	inspectSliceArray([]int{1, 2, 3})
	inspectSliceArray([3]int{4, 5, 6})
}
