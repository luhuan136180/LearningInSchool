package main

import "fmt"
func main(){
	/*//第一题
	var n1 int32
	var n2 int32
	fmt.Println("请输入n1：")
	fmt.Scanln(&n1)
	fmt.Println("请输入n2：")
	fmt.Scanln(&n2)



	if (n1 + n2) > 50 {
		fmt.Println("hello,world")
	}else{
		fmt.Println("两数之和小于50")
	}*/
	//第二题
	var n3 float32 = 11.0
	var n4 float32 = 17.0
	if n3 >10.0 && n4 <20.0{
		fmt.Println("N3+n4=",n3+n4)
	}
	//第三题
	var num1 int32 = 10
	var num2 int32 = 5
	if(num1+num2)%3 ==0 && (num2+num1)%5 == 0{
		fmt.Println("又能被三整除又能被5整除")
	}
	//第四题
	var year int =2019
	if(year % 4 == 0 && year % 100 != 0 || year%400 == 0){
		fmt.Println(year,"是闰年")
	}
}
