package processes

import (
	"ChatDemo/common/message"
	"ChatDemo/server/model"
	"ChatDemo/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加字段，用于判断conn是属于哪个用户的
	UserID int
}

// NotifyOthersProcess 通知所有在线的用户
func (user *UserProcess) NotifyOthersProcess(UserID int, UserStatus int) {

	if UserStatus == message.UserOnline {
		for id, v := range userMgr.onlineUsers {
			if id == UserID {
				continue
			}

			v.NotifyOthersOnline(UserID)
		}
	} else if UserStatus == message.UserOffline {
		for id, v := range userMgr.onlineUsers {
			if id == UserID {
				continue
			}

			//通知下线信息
			v.NotifyOthersOutline(UserID)
		}
	}
}

// NotifyOthersOutline 下线通知
func (user *UserProcess) NotifyOthersOutline(UserID int) {
	//1.封装信息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status = message.UserOffline
	notifyUserStatusMes.UserID = UserID

	//2.序列化消息
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyMes) fail =", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(notifyMes) fail =", err)
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn: user.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg(notify) fail =", err)
		return
	}
}

// NotifyOthersOnline 通知别人上线了
func (user *UserProcess) NotifyOthersOnline(UserID int) {
	//1.封装信息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status = message.UserOnline
	notifyUserStatusMes.UserID = UserID

	//2.序列化消息
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyMes) fail =", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(notifyMes) fail =", err)
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn: user.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg(notify) fail =", err)
		return
	}

}

// ServerLoginProcess 专门处理登录信息
func (user *UserProcess) ServerLoginProcess(mes *message.Message) (err error) {
	//1.先从mes 中取出mes.Data, 并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal(mes.Data) fail =", err)
		return
	}

	//声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//声明一个LoginResMes
	var loginResMes message.LoginResMes

	//2.判断账号信息
	//需要将信息到redis去验证
	_, err = model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_PWD {
			fmt.Println("密码错误...")
		} else if err == model.ERROR_USER_NOTEXISTES {
			fmt.Println("用户不存在...")
		} else if err == model.ERROR_USER_EXISTES {
			fmt.Println("用户已存在...")
		}
	} else {
		//登录成功
		loginResMes.Code = 200
		user.UserID = loginMes.UserID
		userMgr.AddOnlineUser(user)
		//开始通知，成功登录即为上线成功
		user.NotifyOthersProcess(loginMes.UserID, message.UserOnline)
		//将当前在线用户的id放入loginResMes.UsersID
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}

	}

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail =", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对Res序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail =", err)
		return
	}

	//6.发送回复信息
	//使用分层模式MVC，先创建一个transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: user.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg fail =", err)
		return
	}

	return
}

func (user *UserProcess) ServerRegisterProcess(mes *message.Message) (err error) {
	//1.取出mes中的data部分
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail =", err)
		return
	}

	//声明返回信息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	//2.进行redis操作
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTES {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTES.Error()
		}
	} else {
		registerResMes.Code = 200
	}

	return
}

// ServerUsersExit 处理用户离线，调整用户状态
func (user *UserProcess) ServerUsersExit(mes *message.Message) (err error) {
	//1.反序列化信息
	var notifyUserStatusMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail =", err)
		return
	}

	//2.准备在onlineUsers中删除离线用户
	if notifyUserStatusMes.Status == message.UserOffline {
		//准备更新列表
		//先查询该用户是否已经下线，避免重复下线
		_, err = userMgr.GetOnlineUserById(notifyUserStatusMes.UserID)
		if err != nil {
			fmt.Println("GetOnlineUserById fail =", err)
			return
		}
		//删除该在线用户
		userMgr.DeleteOnlineProcess(notifyUserStatusMes.UserID)
	}

	//3.通知其他用户
	user.NotifyOthersProcess(notifyUserStatusMes.UserID, notifyUserStatusMes.Status)

	return
}
