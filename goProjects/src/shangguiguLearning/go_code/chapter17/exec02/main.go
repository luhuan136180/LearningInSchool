package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%d完成了剪发运行，%d -%d=%d\n",name,c.Num1,c.Num2,c.Num1-c.Num2)
}
func TestStruct(a interface{}) {
	//rTpye:=reflect.TypeOf(a)
	rVal:=reflect.ValueOf(a)
	num:=rVal.NumField()
	fmt.Printf("cal共有%d个函数 \n",num)
	for i := 0; i < num; i++ {
		fmt.Printf("field %d:值为=%v \n",i,rVal.Field(i))
	}
	str:="tom"
	var pa []reflect.Value
	pa=append(pa,reflect.ValueOf(str))
	rVal.Method(0).Call(pa)
}
func main() {
	var cal Cal=Cal{
		Num1: 12,
		Num2: 6,
	}
	TestStruct(cal)
}
