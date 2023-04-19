package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestOrder(t *testing.T) {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// 用 age 排序，年龄相等的元素保持原始顺序
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]

	//下面实现排序order by age asc, name desc，如果 age 和 name 都相等则保持原始排序
	sort.SliceStable(family, func(i, j int) bool {
		if family[i].Age != family[j].Age {
			return family[i].Age < family[j].Age
		}
		return strings.Compare(family[i].Name, family[j].Name) == 1
	})

	fmt.Println(family) // [{Eve 2} {David 2} {Alice 23} {Bob 25}]
}
