package main

import "fmt"

func main() {

	// 在 Go 语言中，函数也是一种类型，可以和其他类型一样保存在变量中，
	var f func()
	f = fire
	f() // fire
}
func fire() {
	fmt.Println("fire")
}
