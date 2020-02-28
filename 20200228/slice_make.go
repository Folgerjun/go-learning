package main

import "fmt"

func main() {

	// 如果需要动态地创建一个切片，可以使用 make() 内建函数
	/*
		make( []Type, size, cap )
		其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，
		cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	*/
	b := make([]int, 2)
	c := make([]int, 2, 10)
	fmt.Println(b, c)           // [0 0] [0 0]
	fmt.Println(len(b), len(c)) // 2 2
	/*
		注意：
		使用 make() 函数生成的切片一定发生了内存分配操作，
		但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。
	*/
}
