package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient() {
	// 打印
	fmt.Println("启动客户端。。。")
	// 调用Dial函数：参数需要指定tcp协议，需要指定服务器端的IP+PORT
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("客户端连接失败， err：", err)
		return
	}
	fmt.Println("客户端连接成功，conn：", &conn)
	// 通过客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) // Stdin 代表终端标准输入
	// 从终端读取一行用户输入的信息
	str, err3 := reader.ReadString('\n')
	if err3 != nil {
		fmt.Println("终端输入失败：", err3)
	}
	n, err4 := conn.Write([]byte(str))
	if err4 != nil {
		fmt.Println("连接失败，err:", err4)
	}
	fmt.Printf("终端数据发送成功，共发送%d字节的数据，并退出", n)
}
