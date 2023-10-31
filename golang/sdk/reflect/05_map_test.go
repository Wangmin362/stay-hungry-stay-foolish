package main

import (
	"fmt"
	"reflect"
	"testing"
)

func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)

		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func Test05(t *testing.T) {
	inspectMap(map[uint32]uint32{
		1: 2,
		3: 4,
	})
}
