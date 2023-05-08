package main

import "fmt"

type Person struct {
	Age  int
	Name string
	Addr Address
}

type Address struct {
	Street string
}

func main() {
	a := Person{Age: 18, Name: "davud", Addr: Address{Street: "jiaozi"}}
	b := a

	fmt.Printf("%v, point=%p addrPoint=%p \n", a, &a, &a.Addr)
	fmt.Printf("%v, point=%p addrPoint=%p \n", b, &b, &b.Addr)
	fmt.Println("============================================")
	b.Age = 15
	b.Name = "TOM"
	b.Addr.Street = "dalin"
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}
