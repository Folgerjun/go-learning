package main

import "fmt"

func main() {
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	/*
		fibonacci(1) is: 1
		fibonacci(2) is: 1
		fibonacci(3) is: 2
		fibonacci(4) is: 3
		fibonacci(5) is: 5
		fibonacci(6) is: 8
		fibonacci(7) is: 13
		fibonacci(8) is: 21
		fibonacci(9) is: 34
		fibonacci(10) is: 55
	*/
}

func fibonacci(n int) int {

	if n <= 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
