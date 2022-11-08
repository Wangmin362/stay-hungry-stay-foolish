package main

import (
	"fmt"
	"testing"
)

func TestGoto1(t *testing.T) {
	a := 44
	fmt.Println("aaa")
	if a == 45 {
		goto error
	}
	fmt.Println("bbb")
error:
	fmt.Println("ccc")
	fmt.Println("ddd")
}
