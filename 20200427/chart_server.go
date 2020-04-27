package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp 模拟聊天 服务端
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// 广播
	go broadcaster()

	for { // 循环监听
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 处理每一个客户端 单独开线程
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client) // 连接消息通道
	leaving  = make(chan client) // 离开消息通道
	messages = make(chan string) // 消息通道
)

func broadcaster() {
	// 存储所有连接者
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering: // 连接
			clients[cli] = true

		case cli := <-leaving: // 离开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "欢迎" + who
	messages <- who + " 上线"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() { // 捕获输入
		messages <- who + ":" + input.Text()
	}
	// 这里忽略 input.Err() 中可能的错误
	leaving <- ch
	messages <- who + "下线"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 忽略网络层面的错误
	}
}
