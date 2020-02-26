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
}
