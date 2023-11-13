package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Cat struct {
	Name string
}

func Test02(t *testing.T) {
	var f float64 = 3.5
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	c := Cat{Name: "kitty"}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2.String())
}

type Deployment struct {
	Name string
}

func TestType(t *testing.T) {
	d := &Deployment{}
	ret := reflect.TypeOf(d)
	fmt.Println(ret.Name()) // 返回空
	ret = ret.Elem()
	fmt.Println(ret.Name()) // 返回：Deployment
}

func TestType003(ttt *testing.T) {
	var x float64 = 1.2345

	fmt.Println("==TypeOf==")
	t := reflect.TypeOf(x)
	fmt.Println("type:", t) // 说明通过Type实现的String方法的返回值就是类型
	fmt.Println("kind:", t.Kind())

	fmt.Println("==ValueOf==")
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())   // Type返回对象的底层类型
	fmt.Println("kind:", v.Kind())   // Kind返回当前类型
	fmt.Println("value:", v.Float()) // 获取值
	fmt.Println(v.Interface())       // 获取值
	fmt.Printf("value is %5.2e\n", v.Interface())

	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Println("===kind===")
	type MyInt int
	var m MyInt = 5
	v = reflect.ValueOf(m)
	fmt.Println("kind:", v.Kind()) // Kind() 返回底层的类型 int
	fmt.Println("type:", v.Type()) // Type() 返回类型 MyInt
}

func TestType005(t *testing.T) {
	type student struct {
		Name string `json:"name"`
		Age  int    `json:"age" id:"1"`
	}

	stu := student{
		Name: "hangmeimei",
		Age:  15,
	}

	valueOfStu := reflect.ValueOf(stu)
	// 获取struct字段数量
	fmt.Println("NumFields: ", valueOfStu.NumField())
	// 获取字段 Name 的值
	fmt.Println("Name value: ", valueOfStu.Field(0).String(), ", ", valueOfStu.FieldByName("Name").String())
	// 字段类型
	fmt.Println("Name type: ", valueOfStu.Field(0).Type())

	typeOfStu := reflect.TypeOf(stu)
	for i := 0; i < typeOfStu.NumField(); i++ {
		// 获取字段名
		name := typeOfStu.Field(i).Name
		fmt.Println("Field Name: ", name)

		// 获取tag
		if fieldName, ok := typeOfStu.FieldByName(name); ok {
			tag := fieldName.Tag

			fmt.Println("tag-", tag, ", ", "json:", tag.Get("json"), ", id", tag.Get("id"))
		}
	}
}
