package main

import "fmt"
/*

1)使用 map[string]map[string]sting 的map类型
2)key: 表示用户名，是唯一的，不可以重复
3)如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息,
（包括昵称nickname 和 密码pwd）。
4)编写一个函数 modifyUser(users map[string]map[string]sting, name string) 完成上述功能
*/
func modifyUser(users map[string]map[string]string, name string){
	if users[name]!=nil{
		//这个用户存在
		users[name]["passwd"]="88888888"

	}else{
		//这个用户不存在
		users[name]=make(map[string]string)
		users[name]["passwd"]="12344"
		users[name]["nicheng"]="nicheng"+name
	}
}
func main() {
	user :=make(map[string]map[string]string,10)
	user["tom"]=make(map[string]string,2)
	user["tom"]["passwd"]="3452234"
	user["tom"]["nicheng"]="miaomao"
	user["alice"]=make(map[string]string,2)
	user["alice"]["passwd"]="123456"
	user["alice"]["nicheng"]="aaaaaaa"
	fmt.Println(user)

	modifyUser(user,"jack" )
	modifyUser(user,"alice" )

	fmt.Println(user)
}
