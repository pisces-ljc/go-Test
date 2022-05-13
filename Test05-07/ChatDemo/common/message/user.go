package message

// User 定义用户结构体
type User struct {
	//为了序列化和反序列化成功，必须保证用户信息的json字符串的key和结构体字段对应的tag名字一致
	UserID     int    `json:"userID"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}
