package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // 匿名字段
	innerS // 匿名字段
}

func main() {

	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)     // outer.b is: 6
	fmt.Printf("outer.c is: %f\n", outer.c)     // outer.c is: 7.500000
	fmt.Printf("outer.int is: %d\n", outer.int) // outer.int is: 60
	fmt.Printf("outer.in1 is: %d\n", outer.in1) // outer.in1 is: 5
	fmt.Printf("outer.in2 is: %d\n", outer.in2) // outer.in2 is: 10

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 is:%v", outer2) // outer2 is:{6 7.5 60 {5 10}}
}
