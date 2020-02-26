package main

import "fmt"

func main() {

	// 常量生成器 iota
	type Weekday int
	const (
		// const iota = 0
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf(
		`Sunday:%d
Monday:%d
Tuesday:%d
Wednesday:%d
Thursday:%d
Friday:%d
Saturday:%d`, Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 模拟枚举
	var today Weekday = Wednesday
	fmt.Printf("\ntoday is %d\n", today) // today is 3

	/*
		一个 const 声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加，
		类似于电子表格中单元格自动填充的效果，只需要建立好单元格之间的变化关系，拖动右下方的小点就可以自动生成单元格的值。
	*/
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue) // 1 2 4 8
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue) // 1 10 100 1000
}
