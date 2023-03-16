package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestBase(t *testing.T) {
	name := "D:\\Project\\github\\k8s\\controller-runtime\\examples\\scratch-env\\generic1_test.go"
	fmt.Printf("filepath.Base is: %s\n", filepath.Base(name))
	abs, _ := filepath.Abs(name)
	fmt.Printf("filepath.Abs is: %s\n", abs)
	fmt.Printf("filepath.Dir is: %s\n", filepath.Dir(name))
}
