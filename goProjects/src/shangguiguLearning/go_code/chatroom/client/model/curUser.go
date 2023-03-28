package model

import (
	"net"
	"shangguiguLearning/go_code/chatroom/common/message"
)
type CurUser struct {
	Conn net.Conn
	message.User
}
