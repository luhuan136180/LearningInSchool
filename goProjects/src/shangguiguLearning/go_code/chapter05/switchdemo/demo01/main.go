package main

import "fmt"

func main(){
	//var key byte
	//fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	//fmt.Scanf("%c",&key)
	//switch key {
	//	case 'a':
	//		fmt.Println("1")
	//	case 'b':
	//		fmt.Println("2")
	//	case 'c':
	//		fmt.Println("3")
	//	case 'd':
	//		fmt.Println("4")
	//	default:
	//		fmt.Println("输入有误")
	//}



	////switch 后也可以不带表达式，类似 if --else分支来使用。【案例演示】
	//case中可以对范围进行判断
	var score int = 90
	switch {
	case score >90:
		fmt.Println("1")
	case score >=70&&score<=90:
		fmt.Println("2")
	case score >=60&&score <70:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}


	//switch后面可以直接声明一个变量，，分号结束
	switch grade :=90;{
	case grade > 90:
		fmt.Println("2")
	case grade >=70&&grade <=90:
		fmt.Println("1")
	case grade >=60&&grade <70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格~")
	}


	//switch的穿透 fallthrough
	var num int =10
	switch num{
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
		fallthrough
	default:
		fmt.Println("没")
	}
}