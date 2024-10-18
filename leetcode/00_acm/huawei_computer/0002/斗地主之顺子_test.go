package main

import (
	"fmt"
	"testing"
)

func TestShunzi(t *testing.T) {
	//pk := []string{"2", "3", "4", "5", "6", "7", "8", "J", "J", "Q", "K", "A", "2"}
	//res := pooke(pk)
	//fmt.Println(res)

	//pk := []string{"2", "2", "3", "3", "4", "4", "5", "5", "6", "6", "7", "7", "2"}
	//res := pooke(pk)
	//fmt.Println(res)

	pk := []string{"5", "5", "6", "6", "7", "7", "8", "9", "10", "J", "K", "Q", "J", "A"}
	res := pooke(pk)
	fmt.Println(res)
}
