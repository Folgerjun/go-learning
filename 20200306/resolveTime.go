package main

import "fmt"

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func main() {

	// 打印所有返回值
	fmt.Println(resolveTime(8888)) // 0 2 148

	// 自定义获取返回值
	_, hour, minute := resolveTime(8888)
	fmt.Println(hour, minute) // 2 148

}

func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute

	return
}
