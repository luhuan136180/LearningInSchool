package main

import "fmt"

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu.Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say() i",i)

}
type AInterFace interface {//接口会定义好方法的传参，输出等
	Say()
}
type  BInterFace interface {
	Hello()
}
type Monster struct {

}
func (m Monster)Hello(){
	fmt.Println("monster Hello()~")
}
func (m Monster) Say() {
	fmt.Println("monsterSay()~")
}

func main() {
	var stu01 Stu//结构体变量，实现了 Say() 实现了 AInterface
	var a AInterFace=stu01
	a.Say()
	var i integer = 10
	var b AInterFace = i
	b.Say() // integer Say i = 10

	//Monster实现了AInterface 和 BInterface
	//一个自定义类型可以实现多个接口
	var monster Monster
	var a2 AInterFace=monster
	var b2 BInterFace=monster
	a2.Say()
	b2.Hello()
}
