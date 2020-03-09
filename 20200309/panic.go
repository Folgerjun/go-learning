package main

func main() {

	// 手动触发宕机
	panic("crash")

	/*
		panic: crash

		goroutine 1 [running]:
		main.main()
			F:/GoWork/go-learning/20200309/panic.go:4 +0x40
	*/

}
