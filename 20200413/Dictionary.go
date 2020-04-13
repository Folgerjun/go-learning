package main

import "fmt"

// 字典结构
type Dictionary struct {
	data map[interface{}]interface{} // 键值都为 interface{} 类型
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// 设置键值
func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值 如果回调为 false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 清空所有的数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}

	// 初始化
	d.Clear()
	return d
}

func main() {
	// 创建字典实例
	dict := NewDictionary()

	// 添加数据
	dict.Set("a", 30)
	dict.Set("b", 20)
	dict.Set("c", 10)

	// 获取并打印值
	value := dict.Get("b")
	fmt.Println("b:", value) // b: 20

	// 遍历所有的字典元素
	dict.Visit(func(k, v interface{}) bool {
		// 转换
		if v.(int) > 20 {
			fmt.Println(k, "so expensive...")
			return true
		}
		fmt.Println(k, "cheap cheap...")
		return true
	})
}
