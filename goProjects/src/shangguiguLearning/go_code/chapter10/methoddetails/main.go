package main

import "fmt"

type integer int

func (i integer)print() {
	fmt.Println("i=",i)
}
//编写一个方法，可以改变i的值
func(i *integer) change() {
	*i +=1
}

type student struct {
	Name string
	Age int
}
//给*Student实现方法String()
func(stu *student) String() string {
	str := fmt.Sprintf("Name=[%v] ,Age=[%v]",stu.Name,stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	fmt.Println("i=",i)
	i.change()
	fmt.Println("i=",i)

	//定义一个student变量
	stu:=student{
		Name: "smith",
		Age: 82,
	}
	//如果你实现了 *Student 类型的 String方法，就会自动调用
	//不然该语句会输出stu的地址
	fmt.Println(&stu)
}
