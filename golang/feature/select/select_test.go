package _select

import (
	"fmt"
	"testing"
	"time"
)

func TestBaseSelect(t *testing.T) {
	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker1.C:
			fmt.Println("ticker1")
		case <-ticker2.C:
			fmt.Println("ticker2")
		}
	}

}
