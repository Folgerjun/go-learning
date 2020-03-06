package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%d 的阶乘是 %d \n", i, factorial(uint64(i))) // 10 的阶乘是 3628800
}

// 一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，
// 并且 0 的阶乘为 1，自然数 n 的阶乘写作n!，“基斯顿·卡曼”在 1808 年发明了n!这个运算符号。
func factorial(n uint64) uint64 {

	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}
