package main

import (
	"ChatDemo/common/message"
	"ChatDemo/server/processes"
	"ChatDemo/utils"
	"fmt"
	"io"
	"net"
)

type MainProcessor struct {
	Conn net.Conn
}

// ServerProcessMes 根据客户端发送消息的种类不同，决定调用不同的函数来进行处理
func (p *MainProcessor) serverProcessMes(mes *message.Message) (err error) {
	//注：服务器端的conn不能共用，因为归属于不同的人
	switch mes.Type {
	case message.LoginMesType:
		//处理登录操作
		//创建一个userProcess实例
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerLoginProcess(mes)
	case message.RegisterMesType:
		//处理注册操作
		//创建一个userProcess实例
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerRegisterProcess(mes)
	case message.SmsMesType:
		//进行消息发送处理
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.NotifyUserStatusMesType:
		//调整用户离线状态
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerUsersExit(mes)
	case message.PrivateSmsMesType:
		//进行消息发送处理
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendPrivateMes(mes)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}

	return
}

//总控制函数
func (p *MainProcessor) mainProcess() (err error) {
	//读客户端的信息
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已经退出, 服务端也退出！")
				return err
			} else {
				fmt.Printf("conn.Read(bytesLen), err = %v\n", err)
				return err
			}
		}

		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes fail =", err)
			return err
		}
	}

	return
}
