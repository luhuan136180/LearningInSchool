package main

import "fmt"

func main()  {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18{
		fmt.Println("你已经成年，是：",age)
	}else {
		fmt.Println("你的年龄不大于18，为：",age)
	}
}
