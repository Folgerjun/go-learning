package main

import "fmt"

func main() {

	// 定义三个整数的数组
	var a [3]int
	// 打印第一个元素
	fmt.Println(a[0]) // 0
	// 打印最后一个元素
	fmt.Println(a[len(a)-1]) // 0

	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	/*
		0 0
		1 0
		2 0
	*/

	// 只打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	/*
		0
		0
		0
	*/

	// 初始化数组
	var q = [3]int{1, 2, 3}
	fmt.Println(q[2]) // 3
	var r = [3]int{1, 2}
	fmt.Println(r[2]) // 0

	// 如果在数组长度的位置出现“...”省略号，
	// 则表示数组的长度是根据初始化值的个数来计算
	w := [...]int{1, 2, 3}
	fmt.Printf("%T\n", w) // [3]int

	// 比较两个数组是否相等
	/*
		如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，
		我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，
		不能比较两个类型不同的数组，否则程序将无法完成编译。
	*/
	b := [2]int{1, 2}
	c := [...]int{1, 2}
	d := [2]int{1, 3}
	fmt.Println(b == c, b == d, c == d) // true false false

	// 遍历数组
	var sports [3]string
	sports[0] = "basketball"
	sports[1] = "football"
	sports[2] = "table tennis"

	for k, v := range sports {
		fmt.Println(k, v)
	}
	/*
		0 basketball
		1 football
		2 table tennis
	*/
}
