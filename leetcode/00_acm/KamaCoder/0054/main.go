package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scanln(&str)

	res := strings.Builder{}
	for idx := range str {
		if str[idx] >= '0' && str[idx] <= '9' {
			res.WriteString("number")
		} else {
			res.WriteByte(str[idx])
		}
	}
	fmt.Println(res.String())
}
