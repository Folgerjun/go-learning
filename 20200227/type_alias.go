package main

import (
	"fmt"
	"reflect"
)

// 定义商标结构
type Brand struct {
}

// 为商标结构添加 Show() 方法
func (t Brand) Show() {

}

// 为 Brand 定义一个别名 FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 声明变量 a 为车辆类型
	var a Vehicle

	// 指定调用 FakeBrand 的 Show
	a.FakeBrand.Show()

	// 取 a 的类型反射对象
	ta := reflect.TypeOf(a)

	// 遍历 a 的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a 的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName:%v, FieldType:%v\n", f.Name, f.Type.Name())
	}
	/*
		FieldName:FakeBrand, FieldType:Brand
		FieldName:Brand, FieldType:Brand

		FakeBrand 是 Brand 的一个别名，在 Vehicle 中嵌入 FakeBrand 和 Brand 并不意味着嵌入两个 Brand，
		FakeBrand 的类型会以名字的方式保留在 Vehicle 的成员中。
	*/
}
