package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
/*
通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。
后面 3 个参数分别如下：
	参数名称：在命令行输入参数时，使用这个名称。
	参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
	参数说明：使用 -help 时，会出现在说明中。
*/
var mode = flag.String("mode", "", "process mode")

/*
 将这段代码命名为 main.go
 然后执行 go run main.go --mode=fast
*/
func main() {

	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode) // fast
}
