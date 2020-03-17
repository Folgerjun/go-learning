package main

import (
	"fmt"
	"sort"
)

// 声明英雄的分类
type HeroKind int

// 定义 HeroKind 常量 类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// 定义英雄名单的结构
type Hero struct {
	Name string   // 英雄的名字
	Kind HeroKind // 英雄的种类
}

// 将英雄指针的切片定义为 Heros 类型
type Heros []*Hero

// 实现 sort.Interface 接口取元素数量方法
func (s Heros) Len() int {
	return len(s)
}

// 实现 sort.Interface 接口比较元素方法
func (s Heros) Less(i, j int) bool {

	// 如果英雄的分类不一致时，优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}

	// 默认按英雄名字字符升序排序
	return s[i].Name < s[j].Name
}

// 实现 sort.Interface 接口交换元素方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	// 准备英雄列表
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	// 使用 sort 包进行排序
	sort.Sort(heros)

	// 遍历
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
	/*
		&{Name:关羽 Kind:1}
		&{Name:吕布 Kind:1}
		&{Name:李白 Kind:2}
		&{Name:貂蝉 Kind:2}
		&{Name:妲己 Kind:3}
		&{Name:诸葛亮 Kind:3}
	*/
}
