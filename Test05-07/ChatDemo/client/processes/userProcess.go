package processes

import (
	"ChatDemo/client/model"
	"ChatDemo/common/message"
	"ChatDemo/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (up *UserProcess) Login(userID int, userPWD string) (err error) {

	//fmt.Printf("userID = %d, userPWD = %s\n", userID, userPWD)
	//
	//return nil

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}

	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPWD

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//5.把登录信息赋值给了message.Data
	mes.Data = string(data)

	//6.将message序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	//处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg(conn) err =", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal(resMes) fail =", err)
		return
	}

	if loginResMes.Code == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.User = message.User{
			UserID:     userID,
			UserStatus: message.UserOnline,
		}

		//fmt.Println("登录成功！")
		//显示当前在线的用户列表
		fmt.Println("**********************************")
		fmt.Println("当前在线用户列表：")
		fmt.Println("**********************************")
		for _, v := range loginResMes.UsersID {
			//不用显示自己在线的ID
			if v == userID {
				continue
			}

			fmt.Println("用户ID =", v)
			user := &message.User{
				UserID:     userID,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("**********************************")
		//这里需要再启动一个协程
		//该协程保持和服务段的通讯，如果服务器有数据推送给客户端
		//则接收并显示客户端的终端
		/************************************/
		go serverProcessMes(conn)
		/************************************/

		//1. 显示登录成功的菜单...
		for {
			ShowMenu(curUser)
		}

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (up *UserProcess) Register(userID int, userPwd, userName string) (err error) {
	//1.连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial fail =", err)
		return
	}
	defer conn.Close()

	//2.封装信息
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//3.序列化封装的信息
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) fail =", err)
		return
	}

	mes.Data = string(data)

	//序列化发送的信息
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail", err)
		return
	}

	//4.准备发送信息
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail =", err)
		return
	}

	//5.接收服务器的回复信息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg() err =", err)
		return
	}

	//6.反序列化回复信息中的data部分
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal() fail =", err)
		return
	}

	if registerResMes.Code == 200 {
		fmt.Println("注册成功！")
	} else {
		fmt.Println()
	}

	return
}

// ExitSys 系统退出
func (up *UserProcess) ExitSys(curUser model.CurUser) (err error) {
	//1.封装信息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = curUser.User.UserID
	notifyUserStatusMes.Status = message.UserOffline

	//2.序列化发送的信息
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) fail =", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) fail =", err)
		return
	}

	//3.向服务器发送消息
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg fail =", err)
		return
	}

	return
}


