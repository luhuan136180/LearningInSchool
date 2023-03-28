package main

import (
	"fmt"
	"net"
	"shangguiguLearning/go_code/chatroom/serve/model"
)

//func writePkg(conn net.Conn, data []byte) (err error) {
//
//	//先发送一个长度给对方
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	// 发送长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//
//	//发送data本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//	return
//}

////编写一个队登录操作进行操作的函数
//func serverProcessLogin(conn net.Conn,mes *message.Message)(err error)  {
//	//核心代码
//	//先从mes中取出mes。Data，直接反序化成LoginMes
//	var loginMes message.LoginMes
//	err=json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err!=nil{
//		fmt.Println("JSON.Unmarshal fail err=",err)
//		return
//	}
//	//先声明一个resMes
//	var resMes message.Message
//	resMes.Type=message.LoginResMesType
//	//在声明一个LoginResMes,并完成赋值
//	var LoginResMes message.LoginResMes
//	//如果用户id=100，密码=123456，认为合法
//	if loginMes.UserId ==100 && loginMes.UserPwd == "123456"{
//		LoginResMes.Code=200//合法
//
//	}else {
//		//不合法
//		LoginResMes.Code=500//500状态妈，表示该用户不存在
//		LoginResMes.Error="该用户不存在，请注册在删除"
//	}//至此，loginResMes消息体内容填充完毕，开始序列化
//	//3.将loginResMes 序列化
//	data,err:=json.Marshal(LoginResMes)
//	if err!=nil{
//		fmt.Println("JSON.Marshal fail:",err)
//		return
//	}
//	//4.将data赋值给resMes
//	resMes.Data=string(data)
//	//5.将resMes进行序列化，准备发送
//	data,err=json.Marshal(resMes)
//	err=writePkg(conn,data)
//	return
//}
//判断消息类型
//功能：根据客户端的消息种类不同，决定调用相应的函数处理
//func serverProcessMes(conn net.Conn,mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录
//		err=serverProcessLogin(conn,mes)
//	case message.RegisterMesType:
//		//处理注册
//	default:
//		fmt.Println("消息类型不存在，无法处理》。。。")
//	}
//	return
//}

//注释：当返回内容在函数头部被命名之后，
//函数内部return后面可以不接返回的字面量，函数会自动
//识别字面量相同类型相同的值返回
//读取消息的方法

//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	//我们将读取数据包封装成一个函数
//	buf:=make([]byte,8096)
//	fmt.Println("读取客户端发送的数据。。。")
//	//conn.Read 在conn没有被关闭的情况下，才会阻塞
//	//如果客户端关闭了 conn 则，就不会阻塞
//	_,err=conn.Read(buf)
//	if err!=nil{
//		fmt.Println("conn.Read(bytes) err", err)
//		return
//	}
//
//	fmt.Println("读到的buf=",buf[:4])
//	//将【】byte型：buf 转换为一个uint32数据类型
//	var pkgLen uint32
//	pkgLen=binary.BigEndian.Uint32(buf[:4])//获取长度
//
//	//根据pkglen 读取消息内容
//	n,err:=conn.Read(buf[:pkgLen])//接收消息
//	//fmt.Println("buf=",buf)//发过来的时数字串
//	fmt.Printf("n=%d,pkglen=%d \n",n,pkgLen)
//	if uint32(n) != pkgLen||err!=nil{//比对读取的信息长度是否与pkglen相同
//
//		fmt.Println("readpkg body error")
//		return
//	}
//	//对信息的比对完成，现在将信息反序列化，获得信息体message
//	//位置原因，反序化是接收需要用地址（指针方式接收）
//	err=json.Unmarshal(buf[:pkgLen],&mes)//将反序化后的信息写入 mes中
//	if err!=nil{
//		fmt.Println("json.Unmarsha err=",err)
//		return
//	}
//	return
//	//
//}

//处理和客户端的通讯
func Process(conn net.Conn) {
	defer conn.Close() //延缓关闭
	//循环的客户端发送的信息
	//这里调用总控，创建一个processor实例
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}

//编写一个对UserDao初始化任务的函数
func initUserDao() {
	model.MyuserDao = model.NewUserDao(pool)
}

func init() {
	initPool("localhost:6379", 16, 0, 300)
	initUserDao()
}
func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("vet.Listen err = ", err)
		return
	} //执行到这里表示err已经有了，也就是listen也有了，

	//一旦监听成功，就等待客户端来链接服务器
	for { //当前阶段循环写死防止主程序结束关闭
		fmt.Println("等待客户端的链接。。")
		conn, err := listen.Accept() //有其他东西在监听端口发送了请求连接，并拿到了C/S的coon
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦链接成功，则启动一个协程和客户端保持通讯。。
		go Process(conn) //启动携程
	}
}
