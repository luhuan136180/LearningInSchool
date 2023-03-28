package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
//
func test() {
	//defer func(){
	//
	//	if  err:=recover();err != nil{
	//		fmt.Println("TEST()发生错误~")
	//	}
	//}()
	var myMap map[int]string
	myMap=make(map[int]string,10)
	myMap[0]="golang"
	fmt.Println(myMap[0])
}
func main() {
	go sayHello()
	go test()

	for i:=0;i<10;i++{
		fmt.Println("MAIN()ok=",i)
		time.Sleep(time.Second)
	}
}
