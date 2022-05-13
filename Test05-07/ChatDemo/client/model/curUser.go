package model

import (
	"ChatDemo/common/message"
	"net"
)

// CurUser 在客户端需要使用
type CurUser struct {
	Conn net.Conn
	User message.User
}
