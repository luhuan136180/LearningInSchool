package model

import (
	"fmt"
	"testing"
)

func TestUser_AddUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{}
	//调用添加用户的方法
	user.AddUser()
	user.AddUser2()
}
func TestUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	user:=&User{}
	user,err:=user.GetUserById(1)
	if err!=nil{
		fmt.Println("ERR=",err)
	}
	fmt.Println(user)
}
func TestUser_GetUsers(t *testing.T) {
	fmt.Println("查询所有的记录：")
	user:= User{}
	var users []*User
	users,err:=user.GetUsers()
	if err!=nil{
		fmt.Println("err=",err)
	}
	for i,v :=range users{
		fmt.Printf("users[%d]=%v",i,v)
	}
}
