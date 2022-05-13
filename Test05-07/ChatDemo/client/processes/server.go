package processes

import (
	"ChatDemo/client/model"
	"ChatDemo/common/message"
	"ChatDemo/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

// ShowMenu 显示登录成功之后的界面
func ShowMenu(curUser model.CurUser) {

	fmt.Printf("---------用户%d已经登录成功-------------\n", curUser.User.UserID)
	fmt.Println("---------1.显示用户在线列表-----------")
	fmt.Println("---------2.群  聊-----------------")
	fmt.Println("---------3.私	聊-----------------")
	fmt.Println("---------4.退出系统-----------------")
	fmt.Printf("请选择（1-4）:\n")
	var key int

	//smsProcess实例需要反复使用，将其定义在外部，节省需要重复定义的开销
	//smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("***************显示列表信息****************")
		outputOnlineUser()
		fmt.Println("*****************************************")
	case 2:
		//fmt.Println("请输入即将发送的内容：")
		fmt.Println("**************欢迎进入群聊界面！*************")
		//fmt.Scanln(&content)
		groupChatMenu()
		fmt.Println("*****************************************")
		//smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("**************欢迎进入私聊界面！*************")
		fmt.Println("当前在线用户：")
		outputOnlineUser()
		fmt.Println("*****************************************")

		privateChatMenu()
		fmt.Println("*****************************************")
	case 4:
		fmt.Println("退出系统")
		//退出系统
		up := &UserProcess{}
		up.ExitSys(curUser)
		os.Exit(0)
	default:
		fmt.Println("请输入正确的数字！")
	}

}

//处理服务器推送过来的消息
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Printf("客户端%s正在等待读取服务器发送的消息\n", conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err)
			return
		}

		//如果读取到了消息，下一步处理
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线
			//处理
			//1.取出数据
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal(notifyMes) fail =", err)
				return
			}
			//2.把这个用户的信息，状态保存到客户map[int]User中
			updateUsersStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//群发消息
			outputGroupMes(&mes)
		case message.PrivateSmsMesType:
			//个人界面显示私聊信息
			outputPrivateMes(&mes)
		default:
			fmt.Println("返回了未知消息类型")

		}
	}

}

//群聊界面
func groupChatMenu() {
	flag := true
	smsProcess := &SmsProcess{}
	for flag {
		//当检测到输入信息为exit时，返回上级目录，否则就一直停留在群聊界面
		fmt.Println("请输入即将发送的内容：")
		var content string
		//scan不支持读取带空格的字符串
		reader := bufio.NewReader(os.Stdin)
		//读取到回车符为止
		content, _ = reader.ReadString('\n')
		//去回车符号
		content = strings.Trim(content, "\r\n")
		//检测到返回词条，则返回上级目录
		if content == "exit" || content == "quit" {
			return
		}
		//群聊
		smsProcess.SendGroupMes(content)
	}
}

//私聊界面
func privateChatMenu() {
	fmt.Printf("选择私聊的用户ID:")
	var userID int
	fmt.Scanln(&userID)
	smsProcess := &SmsProcess{}
	for {
		fmt.Printf("请输入向用户%d发送的信息:\n", userID)
		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadString('\n')
		content = strings.Trim(content, "\r\n")
		if content == "exit" || content == "quit" {
			return
		}

		//私聊
		smsProcess.SendToSomeOneMes(content, userID)
	}
}
