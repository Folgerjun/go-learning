package main

import (
	"fmt"
	"sync"
	"time"
)

func syncMutex1() {
	var a = 0
	var lock sync.Mutex // 互斥锁
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	// 等待 1s 结束主程序 确保所有协程执行完
	// 主程结束协程也会强制结束
	time.Sleep(time.Second)
	fmt.Println("over...")
}

func syncMutex2() {
	ch := make(chan struct{}, 2)
	var l sync.Mutex // 互斥锁
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1:我会大概锁定 2s")
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine1:我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		time.Sleep(time.Millisecond * 500) // 先让 goroutine1 执行
		fmt.Println("goroutine2:等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2:哈哈，我也解锁了")
		ch <- struct{}{}
	}()
	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch // 阻塞等待
	}

}

func main() {
	//syncMutex1()
	syncMutex2()
}
