package main

import "fmt"

func main() {
	bag := new(Bag)
	Insert(bag, 1001)
	fmt.Println(*bag) // {[1001]}

	Insert(bag, 1002)
	fmt.Println(*bag) // {[1001 1002]}

	bag.Insert(1003)
	fmt.Println(*bag) // {[1001 1002 1003]}
}

type Bag struct {
	items []int
}

// 将一个物品放入背包的过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}
