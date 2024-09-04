package main

import (
	"fmt"
	"sort"
)

// https://www.nowcoder.com/practice/de044e89123f4a7482bd2b214a685201

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		m[a] += b
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%d %d\n", key, m[key])
	}

}
