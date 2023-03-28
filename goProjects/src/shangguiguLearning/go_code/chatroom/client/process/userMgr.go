package process

import (
	"fmt"
	"shangguiguLearning/go_code/chatroom/client/model"
	"shangguiguLearning/go_code/chatroom/common/message"
)
//客户端要维护的map
var onlineUsers map[int]*message.User =make(map[int]*message.User,20)
var curUser model.CurUser
//在客户端显示当前在线的用户
func outputOnlineUser()  {
	fmt.Println("当前在线用户列表:")
	for id, _:= range onlineUsers {
		//如果不显示自己.
		fmt.Println("用户id:\t", id)
	}
}
//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes)  {

	//适当优化
	user,ok:= onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId]=user

	outputOnlineUser()
}