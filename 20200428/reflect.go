package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	// 变量的类型名 变量的种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // int int

	// 声明一个空结构体
	type cat struct {
	}
	// 创建 cat 实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类 | Go语言的反射中对所有指针变量的种类都是 ptr，但需要注意的是，指针变量的类型名称是空，不是 *cat。
	fmt.Printf("name:'%v' kinf:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kinf:'ptr'
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name:'%v', element kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // element name:'cat', element kind:'struct'
}
