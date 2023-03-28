package main

import (
	"fmt"
	"os"
	"shangguiguLearning/go_code/chatroom/client/process"
)
var userID int
var userPwd string
var userName string
func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入id号：")
			fmt.Scanln(&userID)
			fmt.Println("请输入密码")
			fmt.Scanln(&userPwd)
			//loop = false
			up:=&process.UserProcess{}
			up.Login(userID,userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入id号：")
			fmt.Scanln(&userID)
			fmt.Println("请输入密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入账户名")
			fmt.Scanln(&userName)
			up:=&process.UserProcess{}
			up.Register(userID,userPwd,userName)

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	//if key==1{
	//	//进入登录chengxu
	//	fmt.Println("请输入id号：")
	//	fmt.Scanln(&userID)
	//	fmt.Println("请输入密码")
	//	fmt.Scanln(&userPwd)
	//	//
	//
	//	//err:= client.login(userID,userPwd)
	//	//if err!=nil{
	//	//	fmt.Println("LOGIN ERR=",err)
	//	//	fmt.Println("登录失败")
	//	//}else{
	//	//	fmt.Println("登录成功")
	//	//}
	//}else if key ==2{
	//	fmt.Println("进行用户注册的逻辑...")
	//}

}