package main

import (
	"fmt"
)

func main() {
	key:=""
	loop:=true
	balance :=1000.0
	money :=0.0
	note:=""
	//定义个变量，记录是否有收支的行为
	flag:=false
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details:="收支\t账户金额\t收支金额\t说   明"

	for{
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&key)


		switch key {
			case "1":
				fmt.Println("-----------------当前收支明细记录-----------------")
				if flag {
					fmt.Println(details)
				}else{
					fmt.Println("目前没有收支明细")
				}
			case "2":
				fmt.Println("本次收入的金额为：")
				fmt.Scanln(&money)
				fmt.Println("本次的收入明细:")
				fmt.Scanln(&note)
				//将这个收入情况，拼接到details变量(加入收支明细zhong)
				//收入    11000           1000            有人发红包
				details+=fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
				flag=true

			case "3":
				fmt.Println("本次支出金额为：")
				fmt.Scanln(&money)
				if money > balance{
					fmt.Println("余额不足，")
					break
				}
				balance -=money
				fmt.Println("本次支出说明:")
				fmt.Scanln(&note)
				details+=fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
				flag= true
			case "4":
				fmt.Println("你确定要退出嘛？ y/n")
				choice:=""
				for{
					fmt.Scanln(&choice)
					if choice=="y"||choice=="no"{
						break
					}
					fmt.Println("你输入的有误")
				}
				if choice=="y"{
					loop = false
				}

			default:
				fmt.Println("请输入正确的选项..")

		}
		if !loop {
			break
		}

	}
	fmt.Println("你退出家庭记账软件的使用...")
}
