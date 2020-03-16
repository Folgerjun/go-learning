package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// 车
type Car struct {
	Wheel
	Engine
}

func main() {
	c := Car{
		Wheel: Wheel{
			Size: 18,
		},
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	// %+v	打印结构体时，会添加字段名
	fmt.Printf("%+v\n", c)
	// {Wheel:{Size:18} Engine:{Power:143 Type:1.4T}}
}
