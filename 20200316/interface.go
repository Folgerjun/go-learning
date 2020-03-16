package main

import "fmt"

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构 用于实现 DataWriter
type file struct {
}

// 实现 DataWriter 接口的 WriteData 方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {

	f := new(file)
	var writer DataWriter

	// 将接口赋值 f，也就是 *file 类型
	writer = f

	// 使用 DataWriter 接口进行数据写入
	writer.WriteData("data")
	/*
		WriteData: data
	*/
}
