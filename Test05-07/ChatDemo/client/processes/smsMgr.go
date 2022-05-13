package processes

import (
	"ChatDemo/common/message"
	"encoding/json"
	"fmt"
)

//输出群聊消息
func outputGroupMes(mes *message.Message) {
	//显示
	//1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal(mes.data) fail =", err)
		return
	}

	//显示信息
	info, _ := fmt.Printf("用户id:\t%d对大家说\t%s\n", smsMes.User.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

//输出私聊信息
func outputPrivateMes(mes *message.Message) {
	//1.反序列化
	var privateSmsMes message.PrivateSmsMes
	err := json.Unmarshal([]byte(mes.Data), &privateSmsMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail =", err)
		return
	}

	//显示信息
	fmt.Println("********有私聊信息!**********")
	fmt.Printf("用户\t%d向你发送消息：%s\n", privateSmsMes.Poster.UserID, privateSmsMes.Content)
	fmt.Println("***************************")
}
