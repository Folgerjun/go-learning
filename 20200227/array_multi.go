package main

import "fmt"

func main() {

	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array) // [[0 0] [20 0] [0 0] [0 41]]

	// 声明两个二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为array2的每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2
	fmt.Println(array1) // [[10 20] [30 40]]

	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array3 = array1[1]
	// 将数组中指定的整型值复制到新的整型变量里
	var value = array1[1][0]
	fmt.Println(array3) // [30 40]
	fmt.Println(value)  // 30
}
