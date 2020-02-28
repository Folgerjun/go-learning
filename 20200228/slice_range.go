package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素
	for index, value := range slice {
		fmt.Printf("Index:%d Value:%d\n", index, value)
	}
	/*
		Index:0 Value:10
		Index:1 Value:20
		Index:2 Value:30
		Index:3 Value:40
	*/

	// 需要强调的是，range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
	/*
		Value: 10 Value-Addr: C00000A138 ElemAddr: C00000E3C0
		Value: 20 Value-Addr: C00000A138 ElemAddr: C00000E3C8
		Value: 30 Value-Addr: C00000A138 ElemAddr: C00000E3D0
		Value: 40 Value-Addr: C00000A138 ElemAddr: C00000E3D8

		因为迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的，
		要想获取每个元素的地址，需要使用切片变量和索引值（例如上面代码中的 &slice[index]）。
	*/

}
