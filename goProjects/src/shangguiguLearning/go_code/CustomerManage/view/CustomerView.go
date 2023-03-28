package view

import (
	//"go_code/CustomerManage/service"
	"fmt"
	"shangguiguLearning/go_code/CustomerManage/service"
)

type customerView struct {
	key  string //接收用户输入...
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段 customerService
	customerService *service.CustomerService
}

func GetcustomerView() *customerView {
	return &customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
}
func (this *customerView) list() {
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].Getinfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}
func (this *customerView) Add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	if this.customerService.Add(service.NewCostomer(name, gender, age, phone, email)) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("添加失败")
	}

}
func (this customerView) Delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")

		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出
func (this *customerView) exit() {

	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Add()
		case "2":
			fmt.Println("修改")
		case "3":
			this.Delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if this.loop == false {
			fmt.Println("退出程序...")
			break
		}
	}
}
func (this *customerView) GetMainMune() {
	this.mainMenu()
}
