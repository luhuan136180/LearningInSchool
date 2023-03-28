package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct{
	Name string`json:"monstter_name"`
	Age int`json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

func testStruct(){
	monster:=Monster{
		Name :"牛魔王",
		Age : 500 ,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}
	//monster序列化
	//Marshal函数返回v的json编码。
	data,err:=json.Marshal(&monster)//func Marshal(v interface{}) ([]byte, error)
	if err!=nil{
		fmt.Printf("序列号错误 err=%v\n", err)
		return
	}
	fmt.Println(monster)
	fmt.Printf("monster序列化后=%v \n",string(data))
}
//将map序列化
func testMap() {
	var a map[string]interface{}
	//使用map，需要make
	a=make(map[string]interface{})
	a["name"]="honghaier"
	a["age"]=20
	a["address"]="hongyadong"
	//将a这个map进行序列化
	//将monster 序列化
	data,err:=json.Marshal(a)
	if err!=nil{
		fmt.Printf("序列化错误，err=%v \n",err)
		return
	}
	fmt.Printf("a map 序列化后=%v \n",string(data))
}
//演示对切片进行序列化, 我们这个切片 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"]="jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice=append(slice,m1)
	//
	//使用map前，需要先make
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice, m2)

	//
	data,err:=json.Marshal(slice)
	if err!=nil{
		fmt.Printf("序列化错误err=%v \n",err)
		return
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))

}
func testFloat64() {
	var num1 float64=241.23
	data,err:=json.Marshal(num1)
	if err!=nil{
		fmt.Printf("序列化错误err=%v \n",err)
		return
	}
	fmt.Printf("num1 序列化后=%v\n", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
