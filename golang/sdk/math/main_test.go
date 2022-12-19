package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCeil(t *testing.T) {
	fmt.Println(math.Ceil(5))
	fmt.Println(math.Ceil(5.222))
	fmt.Println(math.Ceil(5.456))
	fmt.Println(math.Ceil(5.982))
	fmt.Println(math.Ceil(5.0))
	fmt.Println(math.Ceil(4.982))

}
