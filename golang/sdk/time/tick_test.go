package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C

	fmt.Println("ssss")
}
