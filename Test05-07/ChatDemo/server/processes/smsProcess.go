package processes

import (
	"ChatDemo/common/message"
	"ChatDemo/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (smp *SmsProcess) SendGroupMes(mes *message.Message) {
	//1.取出消息
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal(smsMes) fail =", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Unmarshal(mes) fail =", err)
		return
	}

	//2.遍历服务器端的onlineUsers
	for id, up := range userMgr.onlineUsers {
		//过滤自己，避免给自己发送消息
		if id == smsMes.User.UserID {
			continue
		}
		smp.SendToOnlineUsers(data, up.Conn)
	}
}

// SendToOnlineUsers 将消息转发给每个在线用户
func (smp *SmsProcess) SendToOnlineUsers(data []byte, conn net.Conn) {
	//创建一个transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(data) fail =", err)
		return
	}
}

// SendPrivateMes  消息转发给指定用户
func (smp *SmsProcess) SendPrivateMes(mes *message.Message) {
	//1.反序列化消息
	var privateSmsMes message.PrivateSmsMes
	err := json.Unmarshal([]byte(mes.Data), &privateSmsMes)
	if err != nil {
		fmt.Println("json.Unmarshal（mes） fail", err)
		return
	}

	//2.检索在线用户列表，查看是否在线
	up, ok := userMgr.onlineUsers[privateSmsMes.Receiver.UserID]
	if !ok {
		//用户已经下线，将消息存入离线留言中
	}

	//3.准备将消息发送给该用户
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail =", err)
		return
	}

	smp.SendToSomeOneMes(data, up.Conn)

}

// SendToSomeOneMes 转发
func (smp SmsProcess) SendToSomeOneMes(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg fail", err)
		return
	}
}
