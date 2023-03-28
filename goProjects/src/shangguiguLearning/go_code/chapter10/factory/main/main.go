package main

import (
	"fmt"
	"shangguiguLearning/go_code/chapter10/factory/model"
)
func main() {
	//创建要给Student实例
	var stu01 = model.Student{
		Name: "tom",
		Age: 18,
	}
	fmt.Println(stu01)
	var stu02 =&model.Student{
		Name: "tom~",
		Age: 19,
	}
	fmt.Println(*stu02)
	//当student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu03 = model.NewStudent("jack",15)
	fmt.Println(*stu03)
	var stu04 = model.NewStudent02("jack~",16)
	fmt.Println(stu04)
	//
	fmt.Println("name=", stu03.Name, " age=", stu03.Age)
	//如果 model 包的 student 的结构体的字段 Name 改成 name
	var stu05 = model.NewStudent03("smith",23)
	fmt.Println(*stu05)
	fmt.Println("name=",stu05.GetName(),"\t Age=",stu05.Age)
}