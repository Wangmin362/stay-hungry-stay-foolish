package main

import (
	"fmt"
	"github.com/golang/demo/golang/sdk/unsafe/modify_private_member/pkg"
	"unsafe"
)

func main() {
	p := pkg.Person{}
	fmt.Printf("%+v\n", p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "david"

	addr := (*string)(unsafe.Pointer((uintptr(unsafe.Pointer(&p))) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*addr = "shanghai"

	fmt.Printf("%+v\n", p)
}
