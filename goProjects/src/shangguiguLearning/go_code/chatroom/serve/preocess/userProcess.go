package preocess

import (
	"encoding/json"
	"fmt"
	"net"
	"shangguiguLearning/go_code/chatroom/common/message"
	"shangguiguLearning/go_code/chatroom/serve/model"
	"shangguiguLearning/go_code/chatroom/serve/utils"
)
//
type UserProcess struct {
	//字段？
	Conn net.Conn
	UserId int

}
//这里我们编写通知所有在线的用户的方法
//userId 要通知其它的在线用户，我上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int)  {

	//遍历onlineUsers，然后一个一个的发送NotifyUserStatusMes消息
	for id,up:=range userMgr.onlineUsers{
		//过滤掉自己
		if id==userId{
			continue
		}
		//开始通知【单独的写一个方法】
		up.NotifyMeOnline(userId)
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type= message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status= message.UserOnline
	notifyUserStatusMes.UserId=userId

	//将notifyUserStatusMes序列化
	data,err:=json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给 mes.Data
	mes.Data=string(data)
	//对mes再次序列化，准备发送.
	data,err=json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送,创建我们Transfer实例，发送
	tf:=&utils.Transfer{
		Conn: this.Conn,
	}
	err =tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

//
func (this *UserProcess) ServerProcessRegis(mes *message.Message) (err error) {
	//1.先从mes 中取出 mes.Data ，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err =json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type= message.RegisterResMesType
	var registerResMes message.RegisterResMes
	//我们需要到redis数据库去完成注册.
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyuserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}
	data,err:=json.Marshal(registerResMes)
	if err!=nil{
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4. 将data 赋值给 resMes
	resMes.Data = string(data)
	//5. 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6. 发送data, 我们将其封装到writePkg函数
	//因为使用分层模式(mvc), 我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//编写一个对登录操作进行操作的函数
func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error)  {
	//核心代码
	//先从mes中取出mes。Data，直接反序化成LoginMes
	var loginMes message.LoginMes
	err=json.Unmarshal([]byte(mes.Data),&loginMes)
	if err!=nil{
		fmt.Println("JSON.Unmarshal fail err=",err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type= message.LoginResMesType
	//在声明一个LoginResMes,并完成赋值
	var LoginResMes message.LoginResMes
	//我们需要reids数据库去完成验证
	//1.使用model。myuserdao到redis去验证
	user ,err:= model.MyuserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err!=nil{
		if err== model.ERROR_USER_NOTEXISTS {
			LoginResMes.Error=err.Error()
		}else if err== model.ERROR_USER_PWD {
			LoginResMes.Code=403
			LoginResMes.Error=err.Error()
		}else {
			LoginResMes.Code=505
			LoginResMes.Error="服务器内部问题"
		}
		//LoginResMes.Code=500
		//LoginResMes.Error="该用户不存在，请注册后再使用"
	}else {
		LoginResMes.Code=200
		//这里，因为用户登录成功，我们就把该登录成功的用放入到userMgr中
		//将登录成功的用户的userId 赋给 this
		this.UserId =loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其它的在线用户， 我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id,_ := range userMgr.onlineUsers{
			LoginResMes.UsersId=append(LoginResMes.UsersId,id)//将数据导入客户端，
		}
		fmt.Println(user,"登陆成功")
	}
	////如果用户id=100，密码=123456，认为合法
	//if loginMes.UserId ==100 && loginMes.UserPwd == "123456"{
	//	LoginResMes.Code=200//合法
	//
	//}else {
	//	//不合法
	//	LoginResMes.Code=500//500状态妈，表示该用户不存在
	//	LoginResMes.Error="该用户不存在，请注册在删除"
	//}//至此，loginResMes消息体内容填充完毕，开始序列化
	//3.将loginResMes 序列化
	data,err:=json.Marshal(LoginResMes)
	if err!=nil{
		fmt.Println("JSON.Marshal fail:",err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data=string(data)
	//5.将resMes进行序列化，准备发送
	data,err=json.Marshal(resMes)
	//6.发送data，将其封装到writepkg函数中
	//因为使用分层式（mvc），先创建transfer实例，然后读取
	tf:=&utils.Transfer{
		Conn: this.Conn,
	}
	err=tf.WritePkg(data)
	return
}






