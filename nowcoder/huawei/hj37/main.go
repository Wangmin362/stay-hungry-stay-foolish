package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	n1 := 1
	n2 := 1
	fn := 0
	for i := 2; i < n; i++ {
		fn := n1 + n2
		n1 = n2
		n2 = fn
	}
	fmt.Println(fn)
}
