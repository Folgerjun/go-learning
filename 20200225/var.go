package main

import "fmt"

func main() {
	var a = 7
	fmt.Println("a:", a) // 7

	a, b := 4, 5
	fmt.Printf("a:%d b:%d \n", a, b) // a:4 b:5

	// 元素值互换
	a, b = b, a
	fmt.Printf("a:%d b:%d \n", a, b) // a:5 b:4

	// 匿名变量 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
	// 此处若是 := 编译会报错
	a, _ = 1, 2
	_, b = 1, 2
	fmt.Printf("a:%d b:%d \n", a, b) // a:1 b:2
}
