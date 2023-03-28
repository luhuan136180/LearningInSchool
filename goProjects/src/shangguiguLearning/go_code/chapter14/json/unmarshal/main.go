package main

import (
	"encoding/json"
	"fmt"
)
type Monster struct {
	Name string
	Age int
	Birthday string //....
	Sal float64
	Skill string
}
////演示将json字符串，反序列化成struct
func unmarshsalStruct() {
	//说明str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
	str := "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	//
	var monster Monster
	//Unmarshal函数解析json编码的数据并将结果存入v指向的值。
	err:=json.Unmarshal([]byte(str),&monster)//func Unmarshal(data []byte, v interface{}) error{}
	if err!=nil{
		fmt.Printf("unmarshal err=%v \n",err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)

}
//将map进行序列化
func testMap() string {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿~~~~~~"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	//将monster 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	//fmt.Printf("a map 序列化后=%v\n", string(data))
	return string(data)

}
//演示将json字符串，反序列化成map
func unmarshsalMap() {
	str:=testMap()
	var a map[string]interface{}
	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err:=json.Unmarshal([]byte(str),&a)
	if err!=nil{
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}
//演示将json字符串，反序列化成切片
func unmarshalSlice(){
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var slice []map[string]interface{}
	err:=json.Unmarshal([]byte(str),&slice)
	if err!=nil{
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}
func main() {
	unmarshsalStruct()
	unmarshsalMap()
	unmarshalSlice()
}
