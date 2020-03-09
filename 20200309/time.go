package main

import (
	"fmt"
	"time"
)

func main() {

	since()

	sub()

}

func since() {
	start := time.Now() // 获取当前时间
	fmt.Println("start: ", start)

	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func sub() {
	start := time.Now() // 获取当前时间
	fmt.Println("start: ", start)

	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
