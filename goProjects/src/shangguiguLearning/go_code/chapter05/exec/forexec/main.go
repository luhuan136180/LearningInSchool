package main

import "fmt"

func main()  {

	//打印1~100之间所有是9的倍数的整数的个数及总和
	var max int64 =100
	var number int64 =0
	var sum int64 = 0
	var i int64 =1
	for ;i<=max;i++{
		if i%9==0{
			number++
			sum +=i
		}
	}
	fmt.Printf("个数是：%v,sum=%v",number,sum)
	//完成下面的表达式输出 ，6是可变的
	fmt.Println("--------------------------------")
	var n int =9
	for i:=0;i<=n;i++{
		fmt.Printf("%v +%v=%v \n",i,n-i,n)
	}
}
