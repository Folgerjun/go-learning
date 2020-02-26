package main

import "fmt"

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

// 定义 ChipType 类型的方法 String()，返回值为字符串类型。
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func main() {

	// 输出 CPU 的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU) // CPU 1

	/*
		String() 方法的 ChipType 在使用上和普通的常量没有区别。
		当这个类型需要显示为字符串时，Go语言会自动寻找 String() 方法并进行调用
	*/
}
