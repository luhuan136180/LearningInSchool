package main

import "fmt"

type MethodUtils struct {
	
}
//1在方法中打印一个 10*8 的矩形
func(me MethodUtils) print() {
	for i:=1;i<=10;i++{
		for j:=1;j<=8;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//2编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
func(me MethodUtils)print2( n1 int ,n2 int){
	for i:=1;i<=n1;i++{
		for j:=1;j<=n2;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//3编写一个方法算该矩形的面积(可以接收长 len，和宽 width)，
//将其作为方法返回值。在 main
//方法中调用该方法，接收返回的面积值并打印。
func(mu MethodUtils) area(lenth float64,width float64) float64{
	res:= lenth*width
	return  res
}
//4判断一个数是奇数还是偶数
func (mu *MethodUtils)JudgeNum(num int)  {
	if num %2==0{
		fmt.Println(num,"是偶数")
	}else {
		fmt.Println(num,"是奇数")
	}
}
//5根据行、列、字符打印 对应行数和列数的字符，比如：行：3，列：2，字符*,则打印相应的效
//果
func (mu MethodUtils)print3(l int , h int,key string)  {
	for i:=1;i<=l;i++{
		for j:=1;j<=h;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}
//6, 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
type Calcuator struct {
	num1 float64
	num2 float64
}
//实现形式 1：分四个方法完成:
func (c Calcuator) getSum() {
	res:= c.num1+c.num2
	fmt.Println("res=",res)
}
func (c Calcuator) getSub() float64{
	res:=c.num1-c.num2
	return res
}
//实现形式 2：用一个方法搞定
func (c *Calcuator)getRes(operator byte) float64 {
	res:=0.0
	switch operator {
	case '+':
		res=c.num1+c.num2
	case '-':
		res = c.num1 - c.num2
	case '*':
		res = c.num1*c.num2
	case '/':
		res = c.num1/c.num2
	default:
		fmt.Println("运算符输入有误")
	}
	return res
}
func main() {
	//编写结构体(MethodUtils)，编程一个方法，方法不需要参数，
	//在方法中打印一个 10*8 的矩形，
	//在 main 方法中调用该方法。
	var mu MethodUtils
	mu.print()
	fmt.Println()
	//编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
	mu.print2(2,3)
	//编写一个方法算该矩形的面积(可以接收长 len，和宽 width)，
	//将其作为方法返回值。在 main
	//方法中调用该方法，接收返回的面积值并打印。
	res:=mu.area(2,3)
	fmt.Println("res=",res)
	//判断一个数是奇数还是偶数
	mu.JudgeNum(3)
	//根据行、列、字符打印 对应行数和列数的字符，
	//比如：行：3，列：2，字符*,则打印相应的效果
	mu.print3(4,5,"#")
	// 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
	//实现形式 1：分四个方法完成:
	//实现形式 2：用一个方法搞定
	 c:= Calcuator{
		 num1: 20,
		 num2:30,
	 }
	 c.getSum()
	 res2:=c.getSub()
	 fmt.Println("sub=",res2)
	 res3:=c.getRes('-')
	 fmt.Println("res3=",res3)
}
