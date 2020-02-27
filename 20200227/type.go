package main

import "fmt"

// 将 NewInt 定义为 int 类型 | Go 1.9 版本之前
// NewInt 会形成一种新的类型 NewInt 本身依然具备 int 类型的特性
type NewInt int

// 将 int 取一个别名叫 IntAlias | Go 1.9 版本之后
// 使用 IntAlias 与 int 等效
type IntAlias = int

func main() {

	// 将 a 声明为 NewInt 类型
	var a NewInt
	// 查看 a 的类型名
	fmt.Printf("a type: %T\n", a) // a type: main.NewInt

	// 将 a2 声明为 IntAlias 类型
	var a1 IntAlias
	fmt.Printf("a1 type:%T\n", a1) // a1 type:int

	/*
		结果显示 a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，
		a2 类型是 int，IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。
	*/

}
