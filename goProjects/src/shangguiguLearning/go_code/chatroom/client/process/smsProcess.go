package process

import (
	"encoding/json"
	"fmt"
	"shangguiguLearning/go_code/chatroom/client/utils"
	"shangguiguLearning/go_code/chatroom/common/message"
)

type SmsProcess struct {
	
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1创建一个Mes
	var mes message.Message
	mes.Type= message.SmsMesType

	//2创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.UserId= curUser.UserId
	smsMes.Content=content
	smsMes.UserStatus= curUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data=string(data)
	//4. 对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	//5. 将mes发送给服务器。。
	tf := &utils.Transfer{
		Conn : curUser.Conn,
	}
	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}