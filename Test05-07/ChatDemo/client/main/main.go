package main

import (
	"ChatDemo/client/processes"
	"fmt"
)

func main() {
	//loop := true
	var key int
	var userID int
	var userPWD string
	var userName string

	for true {
		fmt.Println("----------多人聊天室---------")
		fmt.Println("          1.登录")
		fmt.Println("          2.注册")
		fmt.Println("          3.退出")
		fmt.Printf("      请选择（1-3）:")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录聊天室!")
			fmt.Printf("账号:")
			fmt.Scanln(&userID)
			fmt.Printf("密码:")
			fmt.Scanln(&userPWD)
			//创建一个用户处理的实例
			up := &processes.UserProcess{}
			err := up.Login(userID, userPWD)
			if err != nil {
				fmt.Println("登录出错， err =", err)
			}
		case 2:
			//1.输入注册信息
			fmt.Println("开始注册！")
			fmt.Printf("账号:")
			fmt.Scanln(&userID)
			fmt.Printf("密码:")
			fmt.Scanln(&userPWD)
			fmt.Print("用户名:")
			fmt.Scanln(&userName)
			//2.调用注册处理
			//创建一个用户处理的实例
			up := &processes.UserProcess{}
			err := up.Register(userID, userPWD, userName)
			if err != nil {
				fmt.Println("register fail =", err)
				return
			}
		case 3:
			//loop = false
		default:
			fmt.Println("请输入正确的选项(1-3)!")
		}
	}

}
