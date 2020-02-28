package main

import "fmt"

func main() {

	// Go 语言的内建函数 append() 可以为切片动态添加元素
	var a []int

	// 追加一个元素
	a = append(a, 1)
	fmt.Println(a) // [1]

	// 追加多个元素，手写解包方式
	a = append(a, 1, 2, 3)
	fmt.Println(a) // [1 1 2 3]

	// 追加一个切片，切片需要解包
	a = append(a, []int{1, 2, 3}...)
	fmt.Println(a) // [1 1 2 3 1 2 3]

	/*
		不过需要注意的是，在使用 append() 函数为切片动态添加元素时，
		如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。
	*/

	// 切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len:%d cap:%d pointer:%p\n", len(numbers), cap(numbers), numbers)
	}
	/*
		len:1 cap:1 pointer:0xc00000a168
		len:2 cap:2 pointer:0xc00000a180
		len:3 cap:4 pointer:0xc00000e400
		len:4 cap:4 pointer:0xc00000e400
		len:5 cap:8 pointer:0xc0000142c0
		len:6 cap:8 pointer:0xc0000142c0
		len:7 cap:8 pointer:0xc0000142c0
		len:8 cap:8 pointer:0xc0000142c0
		len:9 cap:16 pointer:0xc00007a000
		len:10 cap:16 pointer:0xc00007a000
	*/

	// 除了在切片的尾部追加，我们还可以在切片的开头添加元素
	var b = []int{1, 2, 3}

	// 在开头添加一个元素
	b = append([]int{0}, b...)
	fmt.Println(b) // [0 1 2 3]

	// 在开头添加一个切片
	b = append([]int{-3, -2, -1}, b...)
	fmt.Println(b) // [-3 -2 -1 0 1 2 3]

	// 在第 2 个位置插入 9
	b = append(b[:2], append([]int{9}, b[2:]...)...)
	fmt.Println(b) // [-3 -2 9 -1 0 1 2 3]

	// 在第 2 个位置插入切片
	b = append(b[:2], append([]int{5, 5, 5}, b[2:]...)...)
	fmt.Println(b) // [-3 -2 5 5 5 9 -1 0 1 2 3]

	/*
		在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，
		因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。
	*/
}
