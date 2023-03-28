package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string ="tom"
	fs:=reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n",str)
	var num1 int=12
	fs2:=reflect.ValueOf(&num1)
	fmt.Println("FS2=",fs2.Elem().Int())
	fs2.Elem().SetInt(23)
	fmt.Printf("%v\n",num1)
}
