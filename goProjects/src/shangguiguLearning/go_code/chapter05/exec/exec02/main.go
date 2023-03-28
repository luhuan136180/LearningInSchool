package main

import "fmt"

func main()  {
	//var score int
	//fmt.Println("请输入成绩：")
	//fmt.Scanln(&score)
	//
	//if score==100{
	//	fmt.Println("奖励1000")
	//}else if score > 80 &&score <=99{
	//	fmt.Println("奖励100")
	//}else if score > 60 && score <=80{
	//	fmt.Println("奖励10")
	//}else{
	//	fmt.Println("无奖励")
	//}

	////题目2使用陷阱.....只会输出ok1...
	//var n int = 10
	//if n >9 {
	//	fmt.Println("ok1")
	//}else if n>6{
	//	fmt.Println("ok2")
	//}else{
	//	fmt.Println("ok3")
	//}

	//题目3
	//var b bool = true
	//if b == false { 	//如果写成 b = false; 能编译通过吗？如果能，结果是？
	//	fmt.Println("a")
	//} else if b {
	//	fmt.Println("b") // b
	//} else if !b {
	//	fmt.Println("c")//c
	//} else {
	//	fmt.Println("d")
	//}

	//题目4
	//var a float64 = 2.0
	//var b float64 = 2.0
	//var c float64 = 2.0
	//
	//m := b*b -4*a*c
	//if m>0{
	//	x1:=(-b + math.Sqrt(m)) / 2 * a
	//	x2 := (-b - math.Sqrt(m)) / 2 * a
	//	fmt.Printf("x1=%v,x2=%v",x1,x2)
	//}else if m == 0{
	//	x1:=(-b + math.Sqrt(m)) / 2 * a
	//	fmt.Printf("x=%v",x1)
	//}else{
	//	fmt.Println("没有解")
	//
	//}

	//timu5
	var high int32
	var money float32
	var handsome bool
	fmt.Println("请输入身高")
	fmt.Scanln(&high)
	fmt.Println("请输入工资")
	fmt.Scanln(&money)
	fmt.Println("前请输入是否帅")
	fmt.Scanln(&handsome)

	if high > 180 && money >1.0 &&handsome==true{
		fmt.Println("我一定要嫁给他!!")
	}else if high > 180 || money >1.0 || handsome==true{
		fmt.Println("嫁吧，比上不足，比下有余")
	}else {
		fmt.Println("不嫁....")
	}
}
