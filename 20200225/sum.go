package main

import "fmt"

func main() {
	a, b := 3, 4
	sum := sum(a, b)
	fmt.Printf("a:%d b:%d \n", a, b) // a:3 b:4
	fmt.Println("sum:", sum)         // sum: 7
}

// 形参 a,b int 类型  返回 int 类型
func sum(a, b int) int {
	return a + b
}
