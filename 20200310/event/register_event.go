package main

// 实例化一个通过字符串映射函数切片的 map
var eventByName = make(map[string][]func(interface{}))

// 注册事件 提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]
	// 在列表切片中添加函数
	//for _, call := range list {
	//	fmt.Println(reflect.DeepEqual(call, callback))
	//if reflect.TypeOf(call).Kind() == reflect.TypeOf(callback).Kind() {
	//	fmt.Println("同类型事件已经注册！")
	//	return
	//}
	//}
	list = append(list, callback)
	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]
	// 遍历这个事件的所有回调
	for _, callback := range list {
		callback(param)
	}
}
