package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 设置 log 配置
func init() {
	// 设置 log 前缀
	log.SetPrefix("【http test】")
	// 设置 log 信息 时间|文件名：代码所在行数
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// w 表示 response 对象，返回给客户端的内容都在对象里处理
// r 表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往 w 里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
}

// 检查对应端口是否被占用
func checkHTTP() {
	for {
		time.Sleep(time.Second)
		log.Println("Checking if started...")
		resp, err := http.Get("http://localhost:8888")
		if err != nil {
			log.Println("Failed:", err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println("Not OK:", resp.StatusCode)
			continue
		}
		break
	}
	log.Println("SERVER UP AND RUNNING!")
}

func main() {
	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/ 目录）的处理器，访问根目录，就会进入当前目录。
	//http.Handle("/", http.FileServer(http.Dir(".")))

	// 设置路由，如果访问 /，则调用 index 方法
	http.HandleFunc("/", index)
	go checkHTTP()
	/*
		在 go 语言中 nil 是一个经常使用的，重要的预先定义好的标识符。它是许多中类型的零值表示。
		许多新有其他编程语言开发经验的 go 语言开发者都会把nil看作是其他语言中的 null（NULL）。
		这是并不完全正确，因为 go 中的 nil 和其他语言中的 null 有很多不同点。
	*/
	// 监听本机 8080 端口
	log.Println("Starting server...")
	time.Sleep(time.Second * 7)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
