package main

import "fmt"

func main() {

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y) // 2 1
}

// 交换函数
func swap(a, b *int) {

	// 取 a 指针的值 赋给临时变量 t
	t := *a
	// 取 b 指针的值 赋给 a 指针指向的变量
	*a = *b
	// 将 a 指针的值 赋给 b 指针指向的变量
	*b = t
}

/*
	如果交换操作的是指针值 交换是不成功的
	这就像写有两座房子的卡片放在桌上一字摊开，
	交换两座房子的卡片后并不会对两座房子有任何影响。
*/
