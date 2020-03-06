package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	BinarySearch(arr, 0, len(arr)-1, 30)
	fmt.Println("main arr=", arr)
	/*
		找到了，下标为：6
		main arr= [1 2 5 7 15 25 30 36 39 51 67 78 80 82 85 91 92 97]
	*/
}

// 二分查找
// 假设是升序
func BinarySearch(arr []int, leftIndex int, rightIndex int, findVal int) {

	// 判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if arr[middle] > findVal {
		// 要查找的数在 leftIndex 到 middle-1 之间
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		// 要查找的数在 middle+1 到 rightIndex 之间
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}
