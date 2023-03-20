package main

import "fmt"

type Hello interface {
	SayJoke(joke string)
}

type Person struct {
	name string
	age  int
}

func (p Person) SayJoke(joke string) {
	fmt.Println(joke)
}

func main() {

	doJoke := func(hello Hello) {
		joke := "sdfsdf"
		hello.SayJoke(joke)
	}

	tom := Person{"Tom", 23}
	doJoke(tom)
}
