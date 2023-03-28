package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//将用户id的创建交给逻辑层，代码自动添加，实现唯一
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//用于查寻完整信息
//返回用户信息，格式化的字符串
func (this Customer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name,
		this.Gender, this.Age, this.Phone, this.Email)
	return info
}

//
