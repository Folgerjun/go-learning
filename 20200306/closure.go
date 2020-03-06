package main

import (
	"fmt"
)

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {

	// 返回一个闭包
	return func() int {

		// 累加
		value++

		// 返回一个累加值
		return value
	}
}

func main() {

	/*
		被捕获到闭包中的变量让闭包本身拥有了记忆效应，
		闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
	*/

	// 创建一个累加器, 初始值为 1
	accumulator := Accumulate(1)

	// 累加 1 并打印
	fmt.Println(accumulator()) // 2

	fmt.Println(accumulator()) // 3

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator) // 0xc000006028

	// 创建一个累加器, 初始值为 10
	accumulator2 := Accumulate(10)

	// 累加 1 并打印
	fmt.Println(accumulator2()) // 11

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2) // 0xc000006038
}
