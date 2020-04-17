package main

import "fmt"

func sendClosedChan() {
	// 创建一个整型的通道
	ch := make(chan int)

	// 关闭通道
	close(ch)

	// 打印通道的指针 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch)) // ptr:0xc000046060 cap:0 len:0

	// 给关闭的通道发送数据
	ch <- 1 // panic: send on closed channel 报错
}

func receiveClosedChan() {
	// 创建一个整型带两个缓冲的通道
	ch := make(chan int, 2)

	// 给通道放入两个数据
	ch <- 0
	ch <- 1

	// 关闭缓冲
	close(ch)

	// 遍历
	for i := 0; i < cap(ch)+1; i++ {
		// 从通道中取出数据
		v, ok := <-ch
		// 打印取出数据的状态
		fmt.Println(v, ok)
	}
	/*
		0 true
		1 true
		0 false // 0 表示这个通道的默认值，false 表示没有获取成功，因为此时通道已经空了。
	*/
}

func main() {

	sendClosedChan()

	receiveClosedChan()
}
