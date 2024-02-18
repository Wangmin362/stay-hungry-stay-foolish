package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "david", "set name")
	flag.StringVar(&name, "n", "david", "set name")

	flag.Parse()

	fmt.Println(name)
}
