package main

import (
	"fmt"
)

func main() {

	const pi = 3.14159

	// 批量声明多个常量
	const (
		e = 2.7182818
		a = 54321
	)

	/*
		如果是批量声明的常量，
		除了第一个外其它的常量右边的初始化表达式都可以省略，
		如果省略初始化表达式则表示使用前面常量的初始化表达式，
		对应的常量类型也是一样的。
	*/
	const (
		i = 1
		j
		m = 2
		n
	)
	fmt.Printf("i:%d j:%d m:%d n:%d", i, j, m, n) // i:1 j:1 m:2 n:2
}
