package model

import "fmt"

//1) 创建程序,在 model 包中定义 Account 结构体：在 main 函数中体会 Golang 的封装性。
//2) Account 结构体要求具有字段：账号（长度在 6-10 之间）、余额(必须>20)、密码（必须是六
//3) 通过 SetXxx 的方法给 Account 的字段赋值。(同学们自己完成
type account struct {
	accountno string
	pwd string
	balance float64
}
func NewAccount(accoutNo string ,pwd string ,balance float64)*account {
	if len(accoutNo)<6||len(accoutNo)>12{
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd)!=6{
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance<20{
		fmt.Println("余额数目不对")
		return nil
	}
	return &account{
		accountno: accoutNo,
		pwd: pwd,
		balance: balance,
	}
}
func(account *account) SetPwd(pwd string) {
	account.pwd=pwd
}
func(account *account) GetPwd() string {
	return account.pwd
}
//方法
//1. 存款
func (account *account) Deposite(money float64, pwd string)  {

	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.balance += money
	fmt.Println("存款成功~~")

}

//取款
func (account *account) WithDraw(money float64, pwd string)  {

	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0  || money > account.balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.balance -= money
	fmt.Println("取款成功~~")

}

//查询余额
func (account *account) Query(pwd string)  {

	//看下输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v \n", account.accountno, account.balance)

}
