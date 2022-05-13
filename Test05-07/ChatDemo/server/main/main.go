package main

import (
	"ChatDemo/server/model"
	"fmt"
	"net"
)

//处理和客户端之间的通讯
func process1(conn net.Conn) {
	defer conn.Close()
	//调用总控
	processor := MainProcessor{Conn: conn}

	err := processor.mainProcess()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程出错， err =", err)
		return
	}
}

// 完成对UserDao初始化任务
func initUserDao() {
	//pool是全局变量，在redis.go中已经定义
	// 该初始化在redis pool初始化之后
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//服务器启动时就初始化redis连接池
	initPool("127.0.0.1:6379", 16, 0, 300)
	initUserDao()
}

func main() {
	fmt.Println("服务器在8889监听")

	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.listen err =", err)
		return
	}

	defer listen.Close()

	//监听成功
	for {
		fmt.Println("等待客户端来连接！")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err =", err)
			return
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯...
		go process1(conn)
	}
}
