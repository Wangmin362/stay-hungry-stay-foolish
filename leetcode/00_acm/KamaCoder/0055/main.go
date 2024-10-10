package main

import "fmt"

func main() {
	var k int
	var str string
	fmt.Scanln(&k)
	fmt.Scanln(&str)

	k %= len(str)
	fmt.Println(str[len(str)-k:] + str[0:len(str)-k])
}
