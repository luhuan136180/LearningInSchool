package main

import "fmt"

func main() {
	var sum int
	for i:=1;i<100;i++{
		sum+=i
		if sum >20 {
			fmt.Printf("dangqianshu:%v \n",i)
			break
		}
	}

	//实现登录验证，有三次机会，如果用户名为”张无忌” ,密码”888”提示登录成功，
	//否则提示还有几次机会.
	var rest int =3
	var username string
	var passwd int
	for i:=1;i<=rest;i++{
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码:")
		fmt.Scanln(&passwd)
		if username=="张无忌"&&passwd==888{
			fmt.Println("登录成功")
			break
		}else {
			fmt.Printf("还有%v次机会\n",rest-i)
		}
	}
}
