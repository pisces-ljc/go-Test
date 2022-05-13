package processes

import (
	"ChatDemo/client/model"
	"ChatDemo/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers = make(map[int]*message.User, 10)
var curUser model.CurUser //在登录成功之后初始化

func outputOnlineUser() {
	//遍历
	for id, _ := range onlineUsers {
		fmt.Println("用户ID：\t", id)
	}

}

// 处理返回的信息
func updateUsersStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	if notifyUserStatusMes.Status == message.UserOnline {
		user, ok := onlineUsers[notifyUserStatusMes.UserID]
		if !ok {
			//原来没有
			user = &message.User{
				UserID: notifyUserStatusMes.UserID,
			}
		}

		user.UserStatus = notifyUserStatusMes.Status
		onlineUsers[notifyUserStatusMes.UserID] = user
	} else if notifyUserStatusMes.Status == message.UserOffline {
		delete(onlineUsers, notifyUserStatusMes.UserID)
		fmt.Printf("用户%d下线！\n", notifyUserStatusMes.UserID)
	}

	//重新输入在线列表
	outputOnlineUser()
}
