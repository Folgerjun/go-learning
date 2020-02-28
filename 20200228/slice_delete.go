package main

import "fmt"

func main() {

	/*
		Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，
		根据要删除元素的位置有三种情况，分别是从开头位置删除、从中间位置删除和从尾部删除，其中删除切片尾部的元素速度最快
	*/

	// 从开头位置删除
	a := []int{1, 2, 3}
	a = a[1:]      // 删除开头 1 个元素
	fmt.Println(a) // [2 3]

	/*
		也可以不移动数据指针，但是将后面的数据向开头移动，
		可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	*/
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头 1 个元素
	fmt.Println(a)              // [2 3]

	// 从中间位置删除
	// 对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成
	a = []int{1, 2, 3}
	a = append(a[:1], a[2:]...) // 删除下标为 1 的元素
	fmt.Println(a)              // [1 3]

	a = []int{1, 2, 3}
	a = a[:1+copy(a[1:], a[2:])] // 删除下标为 1 的元素
	fmt.Println(a)               // [1 3]

	// 从尾部删除
	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部 1 个元素
	fmt.Println(a)   // [1 2]

	// 删除切片指定位置的元素
	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的位置和之后的位置
	fmt.Println(seq[:index], seq[index+1:]) // [a b] [d e]
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq) // [a b d e]

	/*
		提示
		连续容器的元素删除无论在任何语言中，都要将删除点前后的元素移动到新的位置，随着元素的增加，
		这个过程将会变得极为耗时，因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高的话，
		就需要考虑更换其他的容器了（如双链表等能快速从删除点删除元素）。
	*/
}
