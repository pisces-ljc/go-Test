package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"        //群聊
	PrivateSmsMesType       = "PrivateSmsMes" //私聊
)

//定义用户状态的常量
const (
	UserOnline = iota
	UserOffline
	//UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息

type LoginMes struct {
	UserID   int    `json:"userID"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

// LoginResMes 登录返回信息
type LoginResMes struct {
	Code    int    `json:"code"` //返回状态码	500表示未注册， 200表示登录成功
	UsersID []int  //增加字段，保护用户id的切片
	Error   string `json:"error"` //错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //200表示成功， 400表示用户已经占有
	Error string `json:"error"` //返回错误信息
}

// NotifyUserStatusMes 配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserID int `json:"userID"`
	Status int `json:"status"`
}

// SmsMes 增加一个SmsMes 发送的消息  群聊信息
type SmsMes struct {
	Content string `json:"content"` //内容
	User    User
}

// PrivateSmsMes 私聊信息
type PrivateSmsMes struct {
	Content  string `json:"content"` //内容
	Poster   User   //信息发送者
	Receiver User   //信息接收者
}
