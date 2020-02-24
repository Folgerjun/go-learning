package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数据生产者 chan<- 声明只写通道
func producer(header string, channel chan<- string) {
	// 无限循环，不停生产数据
	for true {
		// 将随机数字符串 格式化后发送到通道 rand.Int31() int32
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		// 等待 1 秒  time.Second 默认 1 秒
		time.Sleep(time.Second)
	}
}

// 数据消费者 <-chan 声明只读通道
func customer(channel <-chan string) {
	// 不停获取数据
	for true {
		// 从通道中取出数据，此处会阻塞直到信道中返回数据 := 表示直接赋值定义
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建 producer() 函数的并发 goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	// 数据消费函数
	customer(channel)
}
