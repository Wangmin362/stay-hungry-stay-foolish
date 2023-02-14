package tool

import (
	"fmt"
	"testing"
)

func TestAscii(t *testing.T) {
	perfix := []string{}
	for i := 65; i <= 90; i++ {
		perfix = append(perfix, string(rune(i)))
	}
	for i := 97; i <= 122; i++ {
		perfix = append(perfix, string(rune(i)))
	}

	fmt.Println(perfix)
	fmt.Println(len(perfix))
}
