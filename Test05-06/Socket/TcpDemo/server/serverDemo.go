package main

import (
	"fmt"
	"net" //网络开发工具包
)

func process(conn net.Conn) {

	//循环接收客户端发送的数据
	//延迟关闭
	defer conn.Close()

	for {
		//创建一个切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write,那么就协程阻塞
		fmt.Printf("等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("远程客户端已经退出！")
			return
		}
		//3.显示客户端发送的内容到服务器终端
		fmt.Println(string(buf[:n]))
	}

}

func main() {

	fmt.Println("服务器开始监听...")
	//todo:1.监听本地端口
	//1.1.使用的网络协议是tcp
	//1.2.监听本地的8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		//监听失败
		fmt.Println("failed to listen...err =", err)
	}
	defer listen.Close() //延时关闭

	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err =", err)
			return
		} else {
			fmt.Println("Accept(), suc conn=", conn, "remote addr =", conn.RemoteAddr().String())
		}

		//准备一个协程为客户端服务
		go process(conn)
	}
	//成功监听
	fmt.Println("listen suc=", listen)

}
