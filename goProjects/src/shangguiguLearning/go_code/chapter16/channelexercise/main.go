package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main() {
	//定义一个存放任意数据类型的管道 3个数据
	//var allChan chan interface{}
	allChan:=make(chan interface{},3)

	allChan<-10
	allChan<-"tom jack"
	cat:=Cat{"xiaohuamao",20}
	allChan<-cat

	<-allChan
	<-allChan
	newCat:=<-allChan
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)
	//使用类型断言
	a:=newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
