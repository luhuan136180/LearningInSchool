package main

import (
	"fmt"
	"reflect"
)
type Monster struct{
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float64`json:"成绩"`
	Sex string
}
//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1+n2
}
//方法， 接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
//方法，显示s的值
func (s Monster) Print(){
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

//
func TestStruct(a interface{}) {
	//获取reflect.Type 类型
	Rtyp:=reflect.TypeOf(a)
	//获取reflect.Value 类型
	Rval:=reflect.ValueOf(a)
	//获取到a对应的类别
	kd :=Rtyp.Kind()
	fmt.Println("KD 的kind",kd)
	//如果传入的不是struct，就退出
	if kd!=reflect.Struct{
		fmt.Println("expecet struct")
		return
	}
	//获取到该结构体有几个字段
	num:=Rval.NumField()
	fmt.Printf("struct has %d fields\n", num) //4
	//变量结构体的所有字段
	for i:=0;i<num;i++{
		fmt.Printf("field %d:值为=%v \n",i,Rval.Field(i))
		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		tagVal:=Rtyp.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal!=""{
			fmt.Printf("field %d:tag为=%v\n",i,tagVal)

		}
	}
	//获取到该结构体有多少个方法
	numOfMethod :=Rval.NumMethod()
	fmt.Printf("STRUCT has %d method \n",numOfMethod)

	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII码）
	Rval.Method(1).Call(nil) //获取到第二个方法。调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value//设置一个切片
	params=append(params,reflect.ValueOf(10))//先将输入的参数转换成reflect.value
	params=append(params,reflect.ValueOf(40))
	//func (v Value) Call(in []Value) []Value
	//它返回函数所有输出结果的Value封装的切片
	//func (v Value) Method(i int) Value
	//返回v持有值类型的第i个方法的已绑定（到v的持有值的）状态的函数形式的Value封装。
	res:=Rval.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
	//func (v Value) Int() int64
	fmt.Println("res=",res[0].Int())//返回结果, 返回的结果是 []reflect.Value*/

}
func main() {
	var a Monster = Monster{
		Name: "hahaha",
		Age: 1020,
		Score: 30.9,
	}
	TestStruct(a)
}