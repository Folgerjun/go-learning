package main

import "fmt"

func main() {

	// Go语言中 goto 语句通过标签进行代码间的无条件跳转，同时 goto 语句在快速跳出循环、避免重复退出上也有一定的帮助，使用 goto 语句能简化一些代码的实现过程。

	// 利用 goto 跳出双层循环
	goto1()

	// 集中处理
	goto2()
}

func goto1() {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Println(y)
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回 避免执行进入标签
	return

breakHere:
	fmt.Println("done")
}

func goto2() {

	//	err := firstCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	err = secondCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	fmt.Println("done")
	//	return
	//
	//onExit:
	//	fmt.Println(err)
	//	exitProcess()
}
