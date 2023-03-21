package main

type Hello interface {
	Hello() string
}

// 函数类型居然可以用来实现某个具体的方法
type SomeFunc func(a, b string) string

func (f *SomeFunc) Hello() string {

	return ""
}

func main() {

}
