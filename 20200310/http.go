package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	client := &http.Client{}

	// 创建一个 http 请求
	req, err := http.NewRequest("POST", "http://www.163.com/", strings.NewReader("key=value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 为标头添加信息  | User-Agent 表明用户的代理特性
	req.Header.Add("User-Agent", "myClient")

	// 开始请求
	resp, err := client.Do(req)

	// 处理请求的错误
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 读取响应的 Body 部分并打印
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	defer resp.Body.Close()

	/*
		<html>
		<head><title>405 Not Allowed</title></head>
		<body bgcolor="white">
		<center><h1>405 Not Allowed</h1></center>
		<hr><center>nginx</center>
		</body>
		</html>
	*/
}
