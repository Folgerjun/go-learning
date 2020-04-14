package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	for {
		// select 不像 switch，后面并不带判断条件，而是直接去查看 case 语句。
		// 每个 case 语句都必须是一个面向 channel 的操作。
		select {
		case ch <- 0:
			fmt.Println("send value 0")
		case ch <- 1:
			fmt.Println("send value 1")
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}
