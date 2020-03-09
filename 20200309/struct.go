package main

import "fmt"

func main() {

	var p Point
	p.X = 10
	p.Y = 20
	fmt.Println(p) // {10 20}

	p1 := new(Point) // p1 的类型为 *T，属于指针
	p1.X = 30
	p1.Y = 40
	fmt.Println(p1) // &{30 40}

	p2 := &Point{} // p2 的类型为 *T，属于指针
	p2.X = 50
	p2.Y = 60
	fmt.Println(p2) // &{50 60}

	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation.child.child.name) // 我

	addr := Address{
		Province:    "四川",
		City:        "成都",
		ZipCode:     610000,
		PhoneNumber: "0",
	}

	// 等效
	//addr := Address{
	//	"四川",
	//	"成都",
	//	610000,
	//	"0",
	//}

	fmt.Println(addr) // {四川 成都 610000 0}

	// 实例化一个匿名结构体
	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printMsg(msg) // &{1024 hello}
	/*
		匿名结构体的类型名是结构体包含字段成员的详细描述，
		匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用。
	*/
}

type Point struct {
	X int
	Y int
}

type People struct {
	name  string
	child *People
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

// 打印消息 传入匿名结构体
func printMsg(msg *struct {
	id   int
	data string
}) {
	// 打印
	fmt.Println(msg)
}
