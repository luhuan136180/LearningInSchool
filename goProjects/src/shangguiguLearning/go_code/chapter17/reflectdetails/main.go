package main

import (
	"fmt"
	"reflect"
)
//通过反射，修改,
// num int 的值
// 修改 student的值
func main() {
	var num int =10
	rVal:=reflect.ValueOf(num)
	// 看看 rVal的Kind是
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	rVal2:=reflect.ValueOf(&num)
	fmt.Printf("rVal kind=%v\n", rVal2.Kind())
	fmt.Println("NUM=",num)
	rVal2.Elem().SetInt(20)
	fmt.Println("rVal2=",rVal2.Elem().Int())
	fmt.Println("NUM=",num)
}


