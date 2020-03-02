package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。

	// 初始化
	myList := list.New()

	// 元素插入尾部
	myList.PushBack("first")

	// 遍历元素 i:=l.Front() 表示初始赋值，只会在一开始执行一次
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	// first
	fmt.Println()

	// 元素插入列表前面
	myList.PushFront("head")
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	// head first
	fmt.Println()

	// 保存元素句柄
	element := myList.PushBack("add")

	// 在 add 之后添加 add_after
	myList.InsertAfter("add_after", element)
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	// head first add add_after
	fmt.Println()

	// 在 add 之前添加 add_before
	myList.InsertBefore("add_before", element)
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	// head first add_before add add_after
	fmt.Println()

	// 删除句柄
	myList.Remove(element)
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	// head first add_before add_after
	fmt.Println()
}
