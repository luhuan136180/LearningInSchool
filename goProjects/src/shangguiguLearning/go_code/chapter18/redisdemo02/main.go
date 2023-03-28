package main

import (
	"fmt"
	"redigo-master/redis"
)

func main() {
	//	//通过go 向redis 写入数据和读取数据
	//	//1. 链接到redis
	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println("redis.Dial err=",err)
	}
	fmt.Println("CONN succ..",conn)
	defer conn.Close() //关闭..
	//2. 通过go 向redis写入数据 string [key-val]
	_,err=conn.Do("hset","user01","name","tomjerry")
	if err!=nil{
		fmt.Println("hSET err=",err)
		return
	}
	_,err=conn.Do("hset","user01","age",10)
	if err!=nil{
		fmt.Println("hSET err=",err)
		return
	}
	//3. 通过go 向redis读取数据 string [key-val]
	//redis有很多自转方法
	r1,err:=redis.String(conn.Do("hget","user01","name"))
	if err!=nil{
		fmt.Println("SET err",err)
		return
	}
	r2,err:=redis.String(conn.Do("hget","user01","age"))
	if err!=nil{
		fmt.Println("SET err",err)
		return
	}
	r3,err:=redis.StringMap(conn.Do("hgetall","user01"))
	//因为返回 r是 interface{}
	//因为 name 对应的值是string ,因此我们需要转换
	//nameString := r.(string)
	fmt.Println("操作ok,name=",r1)
	fmt.Println("age=",r2)
	fmt.Println("hgetall=",r3)

}

