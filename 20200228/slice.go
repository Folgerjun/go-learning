package main

import "fmt"

func main() {

	// slice [开始位置 : 结束位置]
	/*
		语法说明如下：
			slice：表示目标切片对象；
			开始位置：对应目标切片对象的索引；
			结束位置：对应目标切片的结束索引。
	*/
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2]) // [1 2 3] [2]
	/*
		从数组或切片生成新的切片拥有如下特性：
			取出的元素数量为：结束位置 - 开始位置；
			取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
			当缺省开始位置时，表示从连续区域开头到结束位置；
			当缺省结束位置时，表示从开始位置到整个连续区域末尾；
			两者同时缺省时，与切片本身等效；
			两者同时为 0 时，等效于空切片，一般用于切片复位。
	*/

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)                // [] [] []
	fmt.Println(len(strList), len(numList), len(numListEmpty)) // 0 0 0
	// 切片判定空的结果
	// 声明但未使用的切片的默认值是 nil，strList 和 numList 也是 nil，所以和 nil 比较的结果是 true。
	fmt.Println(strList == nil) // true
	fmt.Println(numList == nil) // true
	// numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false
	fmt.Println(numListEmpty == nil) // false
	// 切片是动态结构，只能与 nil 判定相等，不能互相判定相等。

}
