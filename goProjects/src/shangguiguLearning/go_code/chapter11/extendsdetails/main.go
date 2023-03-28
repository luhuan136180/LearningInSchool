package main

import "fmt"

type A struct {
	Name string
	age int
}

func (a *A) Sayok() {
	fmt.Println("A Sayok",a.Name)
}
func (a *A) hello() {
	fmt.Println("A Hello=",a.Name)
}
type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("b SayOk",b.Name)
}

func main() {
	var b2 *B=&B{
		Name: "tom",

	}

	fmt.Println(*b2)
	(*b2).SayOk()
	fmt.Println("-------------------")
	var b B
	b.Name="jack"
	b.A.Name="tom"
	fmt.Printf("b.name=[%v],b.A.name=[%v]c\n",b.Name,b.A.Name)
	b.age=100
	b.SayOk()
	b.A.Sayok()
	b.hello()
}
