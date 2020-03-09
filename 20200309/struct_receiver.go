package main

import "fmt"

// 定义属性结构
type Property struct {
	value int // 属性值
}

// 设置属性值
func (p *Property) SetValue(v int) {
	// 修改 p 的成员变量
	p.value = v
}

// 取属性值
func (p *Property) Value() int {
	return p.value
}

// 定义点结构
type MyPoint struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p MyPoint) Add(other MyPoint) MyPoint {
	// 成员值与参数相加后返回新的结构
	return MyPoint{p.X + other.X, p.Y + other.Y}
}

func main() {
	p := new(Property)
	p.SetValue(100)
	fmt.Println(p.Value()) // 100

	p1 := MyPoint{1, 1}
	p2 := MyPoint{2, 2}
	result := p1.Add(p2)
	fmt.Println(result) // {3 3}
}
