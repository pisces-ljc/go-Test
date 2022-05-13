package processes

import (
	"ChatDemo/common/message"
	"ChatDemo/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// SendGroupMes 发送群聊消息
func (smp *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User.UserID = curUser.User.UserID
	smsMes.User.UserStatus = curUser.User.UserStatus

	//2.序列化信息
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) fail =", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail =", err)
		return
	}

	//3.向服务器发送消息
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes fail =", err)
		return
	}

	return
}

// SendToSomeOneMes 发送私聊信息
func (smp *SmsProcess) SendToSomeOneMes(content string, userID int) {
	//1.封装信息
	var mes message.Message
	mes.Type = message.PrivateSmsMesType

	var privateSmsMes message.PrivateSmsMes
	privateSmsMes.Content = content
	privateSmsMes.Poster.UserID = curUser.User.UserID
	privateSmsMes.Receiver.UserID = userID

	//2.序列化信息
	data, err := json.Marshal(privateSmsMes)
	if err != nil {
		fmt.Println("json.Marshal(privateSmsMes) fail =", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail =", err)
		return
	}
}
