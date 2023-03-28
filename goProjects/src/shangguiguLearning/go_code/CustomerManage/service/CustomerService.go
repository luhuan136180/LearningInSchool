package service

import "shangguiguLearning/go_code/CustomerManage/model"

//该CustomerService， 完成对Customer的操作,包括
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的 id+1
	customerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//调用model层的用户创建
func NewCostomer(name string, gender string, age int, phone string, email string) model.Customer {
	return model.NewCustomer2(name, gender, age, phone, email)

}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户至列表中
func (receiver *CustomerService) Add(customer model.Customer) bool {
	receiver.customerNum++
	customer.Id = receiver.customerNum
	receiver.customers = append(receiver.customers, customer)
	return true
}

//删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindID(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//根据id下标寻找客户在切片中的对应下标
func (this *CustomerService) FindID(id int) int {
	index := -1
	for i, v := range this.customers {
		if v.Id == id {
			//找到
			index = i
		}
	}
	return index
}
