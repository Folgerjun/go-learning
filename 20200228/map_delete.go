package main

import "fmt"

func main() {

	// Go 语言提供了一个内置函数 delete()，用于删除容器内的元素

	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	fmt.Println(scene) // map[brazil:4 china:960 route:66]

	delete(scene, "bra") // 不存在这个键 所以还是原来的数据
	fmt.Println(scene)   // map[brazil:4 china:960 route:66]

	delete(scene, "brazil")

	fmt.Println(scene) // map[china:960 route:66]

	/*
		有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，
		不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	*/
}
