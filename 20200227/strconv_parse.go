package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Parse 系列函数用于将字符串转换为指定类型的值，
	// 其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()

	// ParseBool() 函数用于将字符串转换为 bool 类型的值，
	// 它只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误
	testParseBool()

	// ParseInt() 函数用于返回字符串表示的整数值（可以包含正负号）
	/*
		函数签名如下： func ParseInt(s string, base int, bitSize int) (i int64, err error)
		参数说明：
		base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
		bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
		返回的 err 是 *NumErr 类型的，如果语法有误，err.Error = ErrSyntax，如果结果超出类型范围 err.Error = ErrRange。
	*/
	testParseInt()

	// ParseUint() 函数的功能类似于 ParseInt() 函数，
	// 但 ParseUint() 函数不接受正负号，用于无符号整型
	testParseUnit()

	// ParseFloat() 函数用于将一个表示浮点数的字符串转换为 float 类型
	/*
		函数签名如下: func ParseFloat(s string, bitSize int) (f float64, err error)
		参数说明：
		如果 s 合乎语法规则，函数会返回最为接近 s 表示值的一个浮点数（使用 IEEE754 规范舍入）。
		bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；
		返回值 err 是 *NumErr 类型的，如果语法有误 err.Error=ErrSyntax，如果返回值超出表示范围，返回值 f 为 ±Inf，err.Error= ErrRange。
	*/
	testParseFloat()
}

func testParseBool() {

	str1 := "110"
	boo1, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Printf("str1: %v\n", err) // str1: strconv.ParseBool: parsing "110": invalid syntax
	} else {
		fmt.Println(boo1)
	}
	str2 := "t"
	boo2, err := strconv.ParseBool(str2)
	if err != nil {
		fmt.Printf("str2: %v\n", err)
	} else {
		fmt.Println(boo2) // true
	}
}

func testParseInt() {

	str := "-11"
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num) // -11
	}
}

func testParseUnit() {

	str := "11"
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

func testParseFloat() {

	str := "3.1415926"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num) // 3.1415926
	}
}
