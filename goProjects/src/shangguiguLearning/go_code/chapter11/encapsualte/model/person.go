package model

import "fmt"
type person struct {
	Name string
	age int
	sal float64
}

func NewPerson(name string)*person {
	return &person{
		Name: name,
	}
}
func NewPerson02(name string,age int,sal float64)*person {
	return &person{
		Name: name,
		age: age,
		sal: sal,
	}
}
//为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person)SetAge(age int)  {
	if age<0||age>=150{
		fmt.Println("年龄范围不正确")
	}else {
		p.age=age
	}
}
func(p *person) GetAge() int{
		return p.age
}
func(p *person) SetSal(sal float64)  {
	if sal >=3000&&sal<=300000{
		p.sal=sal
	}else {
		fmt.Println("薪水范围不正确..")
	}
}
func(p *person) GetSal() float64 {
	return p.sal
}
