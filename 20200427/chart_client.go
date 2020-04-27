package main

import (
	"io"
	"log"
	"net"
	"os"
)

// tcp 模拟聊天 客户端
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 忽略错误
		log.Println("done")
		done <- struct{}{} // 发信号
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台 goroutine 完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
