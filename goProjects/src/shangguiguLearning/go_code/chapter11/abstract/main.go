package main

import "fmt"
//定义一个结构体Account
type Account struct {
	AccountNo string
	pwd string
	Balance float64
}
//方法
//1. 存款
func(account Account) Deposite(money float64,pwd string) {
	if pwd !=account.pwd{
		fmt.Println("你输入的密码：")
		return
	}
	//看看存款金额是否正确
	if money <=0{
		fmt.Println("你输入的金额有误")
		return
	}
	account.Balance+=money
	fmt.Println("存款成功~")
}
//取钱
func(account *Account) WithDraw(money float64, pwd string) {
	//看看输入的密码是否正确
	if pwd !=account.pwd{
		fmt.Println("你输入的密码：")
		return
	}
	if money <=0||money>=account.Balance{
		fmt.Println("你输入的金额有误")
		return
	}
	account.Balance-=money
	fmt.Println("qukuan成功~")
}
//查询余额
func (account *Account) Query(pwd string){
	//kan下输入的密码
	if pwd !=account.pwd{
		fmt.Println("你输入的密码：")
		return
	}
	fmt.Printf("你的账号为=%v , 余额为=%v \n",account.AccountNo,account.Balance)
}
func main() {
	account:=Account{
		AccountNo: "gs111111",
		pwd: "666666",
		Balance: 400.0,
	}
	account.Query("666666")
	account.Deposite(200.0, "666666")
	account.Query("666666")
	account.WithDraw(150.0, "666666")
	account.Query("666666")

}
