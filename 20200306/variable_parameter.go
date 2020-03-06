package main

import (
	"bytes"
	"fmt"
)

func main() {

	myfunc(1, 2, 3) // 1 2 3
	fmt.Println()

	var v1 = 1
	var v2 int64 = 234
	var v3 = "hello"
	var v4 float32 = 1.234
	myPrintf(v1, v2, v3, v4)
	/*
		1 is an int value.
		234 is an int64 value.
		hello is a string value.
		1.234 is an unknown type.
	*/

	fmt.Println(printTypeValue(100, "str", true))
	/*
		value: 100 type: int
		value: str type: string
		value: true type: bool
	*/
}

// 指定传入的参数类型
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
}

// 用 interface{} 传递任意类型数据是 Go 语言的惯例用法，使用 interface{} 仍然是类型安全的
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 打印不同类型参数
func printTypeValue(slist ...interface{}) string {

	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer

	// 遍历
	for _, s := range slist {
		// 格式化为字符串
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		// 写字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}
