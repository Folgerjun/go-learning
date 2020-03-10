package main

import "fmt"

type MyInt int

// 为 MyInt 添加 IsZero 方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为 MyInt 添加 Add 方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero()) // true

	b = 1
	fmt.Println(b.Add(2)) // 3
}
