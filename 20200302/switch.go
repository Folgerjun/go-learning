package main

import "fmt"

func main() {

	// 基本写法
	switch1()

	// 一分支多值
	switch2()

	// 分支表达式
	switch3()

	// fallthrough | 会执行下一个 case | 新编写的代码，不建议使用 fallthrough
	switch4()
}

func switch1() {

	a := "hello"
	switch a {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}
	// hello
}

func switch2() {

	a := "num"
	switch a {
	case "num", "sum":
		fmt.Println("mmm")
	default:
		fmt.Println("what?")
	}
	// mmm
}

func switch3() {

	a := 11
	switch {
	case a > 10 && a < 20:
		fmt.Println(a)
	default:
		fmt.Println("what?")
	}
	// 11
}

func switch4() {

	s := "hello"
	switch {
	case s == "hello":
		fmt.Print("hello ")
		fallthrough
	case s != "world":
		fmt.Print("world")
	}
	// hello world
}
