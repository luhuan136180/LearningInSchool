package model

import (
	"encoding/json"
	"fmt"
	"redigo-master/redis"
	"shangguiguLearning/go_code/chatroom/common/message"
)

//在服务器启动后，就初始化一个userdao的实例，把她做成全局的，当需要redis池时，就直接shiyong
var (
	MyuserDao *UserDao
)

//定义一个UserDao 结构体体
//完成对User 结构体的各种操作.
type UserDao struct {
	pool *redis.Pool

}
//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao=&UserDao{
		pool: pool,
	}
	return
}



//思考一下在UserDao 应该提供哪些方法给我们
//1. 根据用户id 返回 一个User实例+err
func (this *UserDao)getUserById(conn redis.Conn,id int) (user *message.User,err error) {
	//fmt.Println("进入查找用户的程序")
	//通过给定id 去 redis查询这个用户
	res,err:=redis.String(conn.Do("Hget","users",id))
	if err!=nil{
		if err==redis.ErrNil{
			fmt.Println("用户不存在")
			err= ERROR_USER_NOTEXISTS
		}else {
			fmt.Println("其他错误")
		}
		return
	}
	//fmt.Println("有这个用户")
	user = &message.User{}
	//这里我们需要把res 反序列化成User实例
	err=json.Unmarshal([]byte(res),user)
	//fmt.Println(user)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	fmt.Println("反序化成功，user=",user)
	return
}
//完成登录的校验 Login
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int ,userPwd string) (user *message.User,err error) {
	//先从UserDao的连接取出一条连接
	conn:=this.pool.Get()
	defer conn.Close()
	//fmt.Println("userId=",userId)
	user,err = this.getUserById(conn,userId)
	if err!=nil {
		return
	}

	if user.UserPwd!=userPwd{
		err= ERROR_USER_PWD
		return
	}
	return

}
func (this *UserDao)Register(user *message.User)(err error){
	conn:=this.pool.Get()
	defer conn.Close()
	_,err=this.getUserById(conn,user.UserId)
	if err ==nil{
		err= ERROR_USER_EXISTS
		return
	}
	//这时，说明id在redis还没有，则可以完成注册
	data,err:=json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_,err = conn.Do("hset","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}