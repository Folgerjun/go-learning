package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

func inject1() {

	// 控制实例的创建
	inj := inject.New()
	// 实参注入
	inj.Map("ffj")
	inj.MapTo("tonglei", (*S1)(nil))
	inj.MapTo("L1", (*S2)(nil))
	inj.Map(24)
	// 函数反转调用
	inj.Invoke(Format)
}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name: %s, company: %s, level: %s, age: %d!\n", name, company, level, age)
}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func reject2() {
	// 创建被注入实例
	s := Staff{}
	// 控制实例的创建
	inj := inject.New()
	// 初始值注入
	inj.Map("ffj")
	inj.MapTo("tonglei", (*S1)(nil))
	inj.MapTo("L1", (*S2)(nil))
	inj.Map(24)
	// 实现对 struct 注入
	inj.Apply(&s)
	// 打印结果
	fmt.Printf("s: %v\n", s)
}

func main() {
	inject1()
	//reject2()
}
