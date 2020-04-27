package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

	// 将时间戳转为时间格式
	timeObj := time.Unix(timestamp1, 0)
	fmt.Println("timeObj:", timeObj)

	// 获取一周中的周几
	weekday := now.Weekday().String()
	fmt.Println("weekday:", weekday)

	// 加一个小时
	later := now.Add(time.Hour)
	fmt.Println("later:", later)

	// 时间格式化
	// 格式化的模板为 Go 的出生时间 2006 年 1 月 2 号 15 点 04 分 Mon Jan
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))    // 24 小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 12 小时制
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// Parse 函数可以解析一个格式化的时间字符串并返回它代表的时间。
	/*
		ParseInLocation 与 Parse 函数类似，但有两个重要的不同之处：
			第一，当缺少时区信息时，Parse 将时间解释为 UTC 时间，而 ParseInLocation 将返回值的 Location 设置为 loc；
			第二，当时间字符串提供了时区偏移量信息时，Parse 会尝试去匹配本地时区，而 ParseInLocation 会去匹配 loc。
	*/
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2020-04-27 11:27:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println("timeObj1:", timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println("timeObj2:", timeObj2)

	// 定时器 定时器的本质上是一个通道（channel）
	ticker := time.Tick(time.Second) // 1s 定时器
	for i := range ticker {
		fmt.Println(i) // 每秒执行
	}
}
