package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {

	/*
		在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，而指针、切片、映射、通道、函数和接口的零值则是 nil。

		nil 是Go语言中一个预定义好的标识符，有过其他编程语言开发经验的开发者也许会把 nil 看作其他语言中的 null（NULL），其实这并不是完全正确的，因为Go语言中的 nil 和其他语言中的 null 有很多不同点。
	*/

	// nil 不是关键字或保留字
	test1()

	// nil 没有默认类型
	test2()

	// 不同类型 nil 的指针是一样的
	test3()

	// nil 作比较
	test4()

	// nil 是 map、slice、pointer、channel、func、interface 的零值
	test5()

	// 不同类型的 nil 值占用的内存大小可能是不一样的
	test6()
}

func test1() {
	var nil = errors.New("my god")
	// 虽然上面的声明语句可以通过编译，但是并不提倡这么做。
	fmt.Println(nil) // my god
}

func test2() {
	fmt.Printf("%T\n", nil) // <nil>
	fmt.Println(nil)        // <nil>
}

func test3() {
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr) // 0x0
	fmt.Printf("%p\n", num) // 0x0
}

func test4() {

	// nil 标志符不能比较
	// fmt.Println(nil == nil) // invalid operation: nil == nil (operator == not defined on nil)

	// 不同类型的 nil 是不能比较的
	//var m map[int]string
	//var ptr *int
	//fmt.Println(m == ptr) // Invalid operation: m == ptr (mismatched types map[int]string and *int)

	// 两个相同类型的 nil 值也可能无法比较
	//var s1 []int
	//var s2 []int
	//fmt.Println(s1 == s2) // Invalid operation: s1 == s2 (operator == not defined on []int)

	var s1 []int
	fmt.Println(s1 == nil) // true
}

func test5() {

	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)   // map[int]string(nil)
	fmt.Printf("%#v\n", ptr) // (*int)(nil)
	fmt.Printf("%#v\n", c)   // (chan int)(nil)
	fmt.Printf("%#v\n", sl)  // []int(nil)
	fmt.Printf("%#v\n", f)   // (func())(nil)
	fmt.Printf("%#v\n", i)   // <nil>
}

func test6() {
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16

	/*
		具体的大小取决于编译器和架构，上面打印的结果是在 64 位架构和标准编译器下完成的，
		对应 32 位的架构的，打印的大小将减半。
	*/
}
