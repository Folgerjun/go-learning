package main

import (
	"fmt"
)

func main() {

	var cat = 1
	var str = "banana"
	// & 操作符来获取变量的内存地址
	fmt.Printf("%p %p\n", &cat, &str) // 0xc00000a0c8 0xc0000321f0

	var house = "Fenghuang Road 21-3"
	// 取地址
	ptr := &house
	// 打印 ptr 类型 | %T 相应值的类型
	fmt.Printf("ptr type: %T\n", ptr) // ptr type: *string
	// 打印 ptr 的指针地址 | %p 十六进制表示，前缀 0x
	fmt.Printf("address: %p\n", ptr) // address: 0xc000032200
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value) // value type: string
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value) // value: Fenghuang Road 21-3

	// 取地址操作符 & 和取值操作符 * 是一对互补操作符
	// & 取出地址，* 根据地址取出地址指向的值
}
