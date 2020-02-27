package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Append 系列函数用于将指定类型转换成字符串后追加到一个切片中，其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint()。
	// Append 系列函数和 Format 系列函数的使用方法类似，只不过是将转换后的结果追加到一个切片中

	// 声明一个 slice
	b10 := []byte("int (base 10):")
	// 将转换为 10 进制的 string，追加到 slice 中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) // int (base 10):-42

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) // int (base 16):-2a
}
