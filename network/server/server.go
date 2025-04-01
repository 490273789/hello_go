package server

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 连接用完一定要关闭
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// 从conn连接中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
	}

}

func StartServer() {
	fmt.Println("启动服务端。。。")
	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("服务器监听失败，err：", err)
		return
	}

	// 监听成功
	// 循环等待客户端的连接
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("客户端连接失败，err：", err2)
		} else {
			fmt.Printf("客户端连接成功！，con=%v, 接收到的信息是：%v \n", conn, conn.RemoteAddr().String())
		}

		// 准备一个协程处理客户端的服务请求
		go process(conn)
	}
}
