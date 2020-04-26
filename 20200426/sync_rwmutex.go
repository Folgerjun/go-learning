package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count int
var rw sync.RWMutex // 读写锁

func syncRWMutex1() {

	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go syncRWMutex1Read(i, ch)
	}

	for i := 0; i < 5; i++ {
		go syncRWMutex1Write(i, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}

func syncRWMutex1Read(n int, ch chan struct{}) {
	rw.RLock() // 读锁
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func syncRWMutex1Write(n int, ch chan struct{}) {
	rw.Lock() // 写锁
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	//rand.Seed(time.Now().UnixNano()) // 加上后每次启动随机数不同 否则与上次是相同的
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

func syncRWMutex2() {

	// 多个同时读
	go syncRWMutex2Read(1)
	go syncRWMutex2Read(2)
	// 写的时候什么都不能干
	go syncRWMutex2Write(3)
	go syncRWMutex2Read(4)
	time.Sleep(time.Second * 5)
}

func syncRWMutex2Read(i int) {
	fmt.Println(i, "read start")
	rw.RLock()
	fmt.Println(i, "reading")
	time.Sleep(time.Second)
	rw.RUnlock()
	fmt.Println(i, "read over")
}

func syncRWMutex2Write(i int) {
	fmt.Println(i, "write start")
	rw.Lock()
	fmt.Println(i, "writing")
	time.Sleep(time.Second)
	rw.Unlock()
	fmt.Println(i, "write over")
}

/*
读写锁的区别在于：
当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；
当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。
*/
func main() {
	//syncRWMutex1()
	syncRWMutex2()
}
