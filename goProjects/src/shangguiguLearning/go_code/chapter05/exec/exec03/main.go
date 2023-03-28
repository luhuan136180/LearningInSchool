package main

import "fmt"

func main()  {
	//题目一
	//var second float64
	//fmt.Println("请输入秒数：")
	//fmt.Scanln(&second)
	//if second < 8.0{
	//	var gender string
	//	fmt.Println("请输入性别")
	//	fmt.Scanln(&gender)
	//	if gender == "男"{
	//		fmt.Println("进入决赛的男子组")
	//	}else{
	//		fmt.Println("进入决赛的女子组")
	//	}
	//}else{
	//	fmt.Println("out...")
	//}

	//题目二
	var age int
	var month int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	if month >4 &&month <10{
		price :=60
		if age>60{
			fmt.Printf("%v月票价 %v年龄%v ",month,price/3,age)
		}else if age >18{
			fmt.Print("%v月票价 %v年龄%v ",month,price,age)
		}else{
			fmt.Print("%v月票价 %v年龄%v ",month,price/2,age)
		}
	}else{
		if age <18||age >60 {
			fmt.Println("淡季儿童和老人  票价 20")
		}else {
			fmt.Println("淡季成人 票价 40")
		}
	}
}
