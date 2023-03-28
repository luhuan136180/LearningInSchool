package main

import (
	"fmt"
	"redigo-master/redis"
)

func main() {
	//进行连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err!=nil{
		fmt.Println("redis.Dial err=",err)
	}
	fmt.Println("COnn suc ...",conn)
	defer conn.Close()//及时关闭
	//2. 通过go 向redis写入数据 string [key-val]
	_,err=conn.Do("hmset","monster1","name","jack","age",200,"skill","eat")
	if err!=nil{
		fmt.Println("hSET err=",err)
		return
	}
	//输出
	r,err:=redis.StringMap(conn.Do("hgetall","monster1"))
	if err!=nil{
		fmt.Println("ERR=",err)
	}
	fmt.Println("R=",r)
}

