package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
	}

	//fmt.Println("conn suc =", conn)
	//1.客户端发送单行数据，然后退出
	//os.stdin : 终端标准输入
	reader := bufio.NewReader(os.Stdin)

	//从终端读取到用户数据之后，准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read str err =", err)
	}

	//再将line 发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.write err =", err)
	}

	fmt.Printf("客户端发送了%d字节, 并退出\n", n)
}
