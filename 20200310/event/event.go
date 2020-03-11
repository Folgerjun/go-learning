package main

import "fmt"

type Actor struct {
}

// 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {

	a := new(Actor)

	// 注册名为 OnSkill 的回调
	RegisterEvent("OnSkill", a.OnEvent)
	// 再次在 OnSkill 上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	// 调用事件 所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)
}
