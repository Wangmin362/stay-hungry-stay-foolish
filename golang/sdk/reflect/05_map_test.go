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

func TestMapReflectDeepEqual(t *testing.T) {

	// 定义两个相等的 map
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}

	// 使用 DeepEqual 比较两个 map 是否相等
	if reflect.DeepEqual(map1, map2) {
		fmt.Println("map1 and map2 are equal")
	} else {
		fmt.Println("map1 and map2 are not equal")
	}

	// 定义两个相等的 map
	map5 := map[string]int{"a": 1, "b": 2, "c": 3}
	map6 := map[string]int{"b": 2, "c": 3, "a": 1}

	// 使用 DeepEqual 比较两个 map 是否相等
	if reflect.DeepEqual(map5, map6) {
		fmt.Println("map5 and map6 are equal")
	} else {
		fmt.Println("map5 and map6 are not equal")
	}

	// 定义两个不相等的 map
	map3 := map[string]int{"a": 1, "b": 2, "c": 3}
	map4 := map[string]int{"a": 1, "b": 2, "c": 4}

	// 使用 DeepEqual 比较两个 map 是否相等
	if reflect.DeepEqual(map3, map4) {
		fmt.Println("map3 and map4 are equal")
	} else {
		fmt.Println("map3 and map4 are not equal")
	}

}
