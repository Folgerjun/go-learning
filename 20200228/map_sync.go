package main

import (
	"fmt"
	"sync"
)

func main() {

	// Go 语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。

	// unsafe
	// unSafe()  会报错：fatal error: concurrent map read and map write

	/*
		sync.Map 有以下特性：
		无须初始化，直接声明即可。
		sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
		使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
	*/

	// 并发安全
	syncSave()
}
func unSafe() {

	m := make(map[int]int)

	go func() {
		// 不停对 map 进行写入
		for {
			m[1] = 1
			fmt.Println(1)
		}
	}()

	go func() {
		// 不停对 map 进行读取
		for {
			_ = m[1]
			fmt.Println(2)
		}
	}()

	// 无限循环, 让并发程序在后台执行
	for true {
		fmt.Println(3)
	}

}

func syncSave() {

	var scene sync.Map

	// 将键值对保存
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 取值
	fmt.Println(scene.Load("london")) // 100 true

	// 删除
	scene.Delete("london")
	fmt.Println(scene.Load("london")) // <nil> false

	// 遍历
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})
	/*
		iterate: greece 97
		iterate: egypt 200
	*/

	/*
		sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，
		sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
	*/

}
