package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os" //
	"shangguiguLearning/go_code/chatroom/client/utils"
	"shangguiguLearning/go_code/chatroom/common/message"
)
//显示登陆成功后的界面。。。
func ShowMenu(){
	fmt.Println("-------恭喜xxx登录成功---------")
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	//因为，我们总会使用到SmsProcess实例，因此我们将其定义在swtich外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d \n",&key)
	switch key {
	case 1:
		outputOnlineUser()
		//fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("你想对大家说的什么:)")
		fmt.Scanln(&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的有误")

	}
}

//和服务器端保持通信
func ProcessServerMes(Conn net.Conn){
	//创建一个transfer实例, 不停的读取服务器发送的消息
	tf:=&utils.Transfer{
		Conn: Conn,
	}
	for{
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes,err:=tf.ReadPkg()//对方不关闭链接，就在读取程序上阻塞
		if err!=nil{
			fmt.Println("tf.ReadPkg err=", err)
			return//退出
		}
		//如果读取到消息，又是下一步处理逻辑
		switch mes.Type{
		case message.NotifyUserStatusMesType: // 有人上线了

			//1. 取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2. 把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
			//处理
		case message.SmsMesType: //有人群发消息
			outputGroupMes(&mes)
		default :
			fmt.Println("服务器端返回了未知的消息类型")
		}
		//fmt.Println("mes=",mes)
	}
}