package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int8
}

func TestFmt(t *testing.T) {
	p := &Person{
		Age:  25,
		Name: "david",
	}

	fmt.Printf("%#v\n", p)  // &main.Person{Name:"david", Age:25}
	fmt.Printf("%#v\n", *p) // main.Person{Name:"david", Age:25}
}
