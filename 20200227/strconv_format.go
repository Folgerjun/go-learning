package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Format 系列函数实现了将给定类型数据格式化为字符串类型的功能，其中包括 FormatBool()、FormatInt()、FormatUint()、FormatFloat()。

	// FormatBool() 函数可以一个 bool 类型的值转换为对应的字符串类型
	testFormatBool()

	// FormatInt() 函数用于将整型数据转换成指定进制并以字符串的形式返回
	/*
		函数签名如下： func FormatInt(i int64, base int) string
		其中，参数 i 必须是 int64 类型，参数 base 必须在 2 到 36 之间，返回结果中会使用小写字母“a”到“z”表示大于 10 的数字。
	*/
	testFormatInt()

	// FormatUint() 函数与 FormatInt() 函数的功能类似，
	//但是参数 i 必须是无符号的 uint64 类型
	testFormatUnit()

	// FormatFloat() 函数用于将浮点数转换为字符串类型
	/*
		函数签名如下： func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		参数说明：
		bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入。
		fmt 表示格式，可以设置为“f”表示 -ddd.dddd、“b”表示 -ddddp±ddd，指数为二进制、“e”表示 -d.dddde±dd 十进制指数、“E”表示 -d.ddddE±dd 十进制指数、“g”表示指数很大时用“e”格式，否则“f”格式、“G”表示指数很大时用“E”格式，否则“f”格式。
		prec 控制精度（排除指数部分）：当参数 fmt 为“f”、“e”、“E”时，它表示小数点后的数字个数；当参数 fmt 为“g”、“G”时，它控制总的数字个数。如果 prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f。
	*/
	testFormatFloat()
}

func testFormatBool() {
	num := true
	str := strconv.FormatBool(num)
	fmt.Printf("type:%T, value:%v\n", str, str) // type:string, value:true
}

func testFormatInt() {

	var num int64 = 100
	str := strconv.FormatInt(num, 16)
	fmt.Printf("type:%T, value:%v\n", str, str) // type:string, value:64
}

func testFormatUnit() {

	var num uint64 = 110
	str := strconv.FormatUint(num, 16)
	fmt.Printf("type:%T, value:%v\n", str, str) // type:string, value:6e
}

func testFormatFloat() {

	var num float64 = 3.1415926
	str := strconv.FormatFloat(num, 'E', -1, 64)
	fmt.Printf("type:%T, value:%v\n", str, str) // type:string, value:3.1415926E+00
}
