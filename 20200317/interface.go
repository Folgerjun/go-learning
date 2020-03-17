package main

import "fmt"

func main() {

	//interfaceValue()

	//interfaceGetValue()

	interfaceValueCompare()
}

// 将值保存到空接口
func interfaceValue() {

	var any interface{}
	any = 1
	fmt.Println(any) // 1

	any = "hello"
	fmt.Println(any) // hello

	any = false
	fmt.Println(any) // false
}

// 取值
func interfaceGetValue() {

	var a int = 1
	var i interface{} = a
	fmt.Println(i) // 1

	//var b int = i // 编译错误 cannot use i (type interface {}) as type int in assignment: need type assertion
	//fmt.Println(b)

	var b int = i.(int)
	fmt.Println(b) // 1
}

// 比较
func interfaceValueCompare() {
	// a保存整型
	var a interface{} = 100
	// b保存字符串
	var b interface{} = "hi"
	// 两个空接口不相等
	fmt.Println(a == b) // false
}
