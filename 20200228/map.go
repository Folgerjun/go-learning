package main

import "fmt"

func main() {

	/*
		map 是引用类型，可以使用如下方式声明：
			var mapname map[keytype]valuetype
		其中：
			mapname 为 map 的变量名。
			keytype 为键类型。
			valuetype 是键对应的值类型。
		提示：[keytype] 和 valuetype 之间允许有空格。
	*/

	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	fmt.Println("mapLit:", mapLit)           // mapLit: map[one:1 two:2]
	fmt.Println("mapCreated:", mapCreated)   // mapCreated: map[]
	fmt.Println("mapAssigned:", mapAssigned) // mapAssigned: map[one:1 two:2]

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])       // Map literal at "one" is: 1
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"]) // Map created at "key2" is: 3.141590
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])      // Map assigned at "two" is: 3
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])       // Map literal at "ten" is: 0

	/*
		和数组不同，map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，但是也可以选择标明 map 的初始容量 capacity，
		格式如下：
			make(map[keytype]valuetype, cap)

		当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，
		即使只是大概知道容量，也最好先标明。
	*/
}
