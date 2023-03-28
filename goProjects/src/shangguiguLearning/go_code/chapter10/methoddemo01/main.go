package main

import (
	"fmt"
)

type Person struct {
	Name string
}
//创建Person的方法
//给Person类型绑定一方法
func(a Person) test() {
	a.Name="jack"
	fmt.Println("a.name=",a.Name)
}

//给Person结构体添加speak 方法,输出  xxx是一个好人
func (p Person)Speak() {
	fmt.Println(p.Name,"是一个好人")
}

//给Person结构体添加jisuan 方法,可以计算从 1+..+1000的结果,
//说明方法体内可以函数一样，进行各种运算
func (p Person)jisuan()  {
	sum:=0
	for i:=0;i<=1000;i++{
		sum+=i
	}
	fmt.Println(sum)
}

//给Person结构体jisuan2 方法,该方法可以接收一个参数n，计算从 1+..+n 的结果
func (p Person)jisuan2(n int){
	res:= 0;
	for i:=0;i<=n;i++{
		res+=i
	}
	fmt.Println(res)
}

//给Person结构体添加getSum方法,可以计算两个数的和，并返回结果
func(p Person) getSum(n1 int,n2 int)int  {
	sum:=n1+n2
	return sum
}

func main() {
	var p Person
	fmt.Println("name=",p.Name)
	p.Name="Tom"
	p.test()
	fmt.Println("name=",p.Name)
	//下面的使用方式都是错误的
	// var dog Dog
	// dog.test()
	// test()
	p.Speak()
	p.jisuan()
	p.jisuan2(10)
	sum:=p.getSum(1,3)
	fmt.Println(sum)
}
