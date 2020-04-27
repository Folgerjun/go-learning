package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
)

func main() {

	// user
	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户 id：", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组 id：", u.Gid)
	// 用户所在的所有的组的 id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)

	// 使用 Notify 方法来监听收到的信号
	// 使用 stop 方法来取消监听
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	//signal.Stop(c) // 不允许继续往 c 中存入内容
	// 阻塞直到有信号接收
	ss := <-c
	fmt.Println("捕获信号：", ss)
}
