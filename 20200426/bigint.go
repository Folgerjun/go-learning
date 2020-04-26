package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM = 1000       // 求第 1000 位斐波那契数列
var fibs [LIM]*big.Int // 使用数组保存计算出来的数列指针

func fibonacci(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}
func main() {
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("数列第 %d 位：%d\n", i+1, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为：%s\n", delta)
}
