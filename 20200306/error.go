package main

import (
	"errors"
	"fmt"
)

// 定义除数为 0 的错误
var errDivisionByZero = errors.New("division by zero")

func main() {

	fmt.Println(div(1, 0)) // 0 division by zero
	fmt.Println(div(8, 3)) // 2 <nil>
}

func div(dividend, divisor int) (int, error) {

	// 判断除数为 0 的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算 返回空错误
	return dividend / divisor, nil
}
