package main

import (
	"fmt"
	"io"
	"net"
	"shangguiguLearning/go_code/chatroom/common/message"
	"shangguiguLearning/go_code/chatroom/serve/preocess"
	"shangguiguLearning/go_code/chatroom/serve/utils"
)
//先创建一个Processor 的结构体体
type Processor struct {
	Conn net.Conn
}
func (this *Processor)serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes=",mes)

	switch mes.Type {
		case message.LoginMesType:
			//处理登录
			//创建一个userPreocess实例
			up:= preocess.UserProcess{
				Conn: this.Conn,
			}
			err=up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
			up:=&preocess.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessRegis(mes)//type:data
		case message.SmsMesType:
		//创建一个SmsProcess指针 实例完成转发群聊消息.
			smsProcess:=&preocess.SmsProcess{}
			smsProcess.SendGrouopMes(mes)
		default:
		fmt.Println("消息类型不存在，无法处理》。。。")
	}
	return err
}

func (this *Processor) Process2()(err error) {
	//循环读取客户端发送的信息
	for{//进入循环，断开连接时退出
		//第一步，接收消息，调用函数
		//创建一个transfer，完成读包的任务
		tf:= utils.Transfer{
			Conn: this.Conn,
		}
		mes,err:=tf.ReadPkg()
		if err!=nil{
			if err==io.EOF{
				fmt.Println("客户端退出，服务器也退出")
				return err//退出循环
			}else{
				fmt.Println("READPKG ERR=",err)
				return err//退出循环
			}
		}
		fmt.Println("MES=",mes)//测试：接收到消息
		//对消息的type进行分析，对不同的消息进行不同处理，此处调用一个判断并分流的函数
		err=this.serverProcessMes(&mes)
		if err!=nil{
			fmt.Println("ServerProcessMes err=",err)
			return err
		}

	}//循环结束,执行conn.close（）
}