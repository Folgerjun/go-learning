package main

import "fmt"

func main() {

	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
	/*
		宕机后要做的事情2
		宕机后要做的事情1
		panic: 宕机

		goroutine 1 [running]:
		main.main()
			F:/GoWork/go-learning/20200309/panic_defer.go:9 +0x147
	*/

	// 当 panic() 触发的宕机发生时，panic() 后面的代码将不会被运行，
	// 但是在 panic() 函数前面已经运行过的 defer 语句依然会在宕机发生时发生作用
}
