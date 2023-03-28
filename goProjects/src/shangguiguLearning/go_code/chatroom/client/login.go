package client

//func login(userId int, userPwd string) (err error) {
//
//	//下一个就要开始定协议..
//	// fmt.Printf(" userId = %d userPwd=%s\n", userId, userPwd)
//
//	// return nil
//
//	//1. 链接到服务器
//	conn, err := net.Dial("tcp", "localhost:8889")
//	if err != nil {
//	fmt.Println("net.Dial err=", err)
//	return
//	}
//	//延时关闭
//	defer conn.Close()
//
//	//2. 准备通过conn发送消息给服务
//	var mes message.Message
//	mes.Type = message.LoginMesType
//	//3. 创建一个LoginMes 结构体
//	var loginMes message.LoginMes
//	loginMes.UserId = userId
//	loginMes.UserPwd = userPwd
//
//	//4. 将loginMes 序列化
//	//序列化会生成一串【】byte
//	dataOfLoginMes, err := json.Marshal(loginMes)
//	if err != nil {
//	fmt.Println("json.Marshal err=", err)
//	return
//	}
//	// 5. 把data转成字符串，赋给 mes.Data字段
//	mes.Data = string(dataOfLoginMes)
//
//	// 6. 将 mes进行序列化化
//	dataOfMessage, err := json.Marshal(mes)
//	//fmt.Println("dataofmessage=",dataOfMessage)
//	if err != nil {
//	fmt.Println("json.Marshal err=", err)
//	return
//	}
//
//	// 7. 到这个时候 data就是我们要发送的消息
//	// 7.1 先把 data的长度发送给服务器
//	// 先获取到 data的长度->转成一个表示长度的byte切片
//	var pkgLen uint32
//	pkgLen = uint32(len(dataOfMessage))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:4],pkgLen)
//	//发送长度
//	n,err:=conn.Write(buf[:4])
//	if n!=4 || err!=nil{
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//	fmt.Printf("客户端，发送消息的长度=%d 内容=%s \n",len(dataOfMessage),string(dataOfMessage))
//
//	//发送消息本身
//	//fmt.Println("dataofmessage=",dataOfMessage)
//	_,err=conn.Write(dataOfMessage)
//	if err!=nil{
//		fmt.Println("conn.Write(DateOfMessage) fail",err)
//		return
//	}
//	//休眠20秒
//	//time.Sleep(20*time.Second)
//	//fmt.Println("休眠了20秒..")
//	//读取传回的消息
//	mes,err= utils.readPkg(conn)
//	fmt.Println("resmes=",mes)
//	if err!=nil{
//		fmt.Println("readpkg(conn),err=",err)
//		return
//	}
//	//将mes的Data部分反序列化成LoginResMes
//	var loginResMes message.LoginResMes
//	err=json.Unmarshal([]byte(mes.Data),&loginResMes)
//	fmt.Println(loginResMes)
//	if loginResMes.Code==200{
//		fmt.Println("登录成功~")
//	}else if loginResMes.Code==500 {
//		fmt.Println(loginResMes.Error)
//
//	}
//	return
//}