package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Cat struct {
	Name string
}

func Test02(t *testing.T) {
	var f float64 = 3.5
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	c := Cat{Name: "kitty"}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2.String())
}
