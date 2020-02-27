package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Itoa()：整型转字符串
	testItoa()

	// Atoi()：字符串转整型
	testAtoi()
}

func testItoa() {
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", str, str) // type:string value:"100"
}

func testAtoi() {
	str1 := "110"
	str2 := "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1) // type:int value:110
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！\n", str2) // s100 转换失败！
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}
}
