package main

import (
	"fmt"
	"sort"
)

func main() {

	// 切片排序 sort.Strings(sceneList)

	// map 遍历 k 值 保存到切片中 再进行遍历
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	var sceneList []string

	for k := range scene {
		sceneList = append(sceneList, k)
	}

	// 排序
	sort.Strings(sceneList)

	fmt.Println(sceneList) // [brazil china route]
}
