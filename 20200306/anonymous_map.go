package main

import (
	"flag"
	"fmt"
)

// 启动参数
// 定义命令行参数 skill，从命令行输入 --skill 可以将 = 后的字符串传入 skillParam 指针变量
var skillParam = flag.String("skill", "", "skill to perform")

func main() {

	// 解析
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
