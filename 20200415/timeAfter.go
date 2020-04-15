package main

import (
	"fmt"
	"time"
)

func main() {

	// 声明一个退出用的通道
	exit := make(chan int)

	// 打印开始
	fmt.Println("start")

	// 过 1 秒后，调用匿名函数
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		// 通知 main() 的 goroutine 已经结束
		exit <- 0
	})

	// 等待结束 阻塞
	<-exit
	fmt.Println("end")
}
