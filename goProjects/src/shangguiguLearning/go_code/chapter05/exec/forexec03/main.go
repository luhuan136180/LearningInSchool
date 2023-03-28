package main

import "fmt"

func main() {
	var hang int
	fmt.Println("请输入行数:")
	fmt.Scanln(&hang)
	for i := 1; i <= hang; i++ {
		for j := hang - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			fmt.Print("*")
		}

		fmt.Println() //换行
	}
	fmt.Println("------------------")
	//5 打印空心金字塔
	for i := 1; i <= hang; i++ {
		for j := hang - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == 2*i-1 || i == hang {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println() //换行
	}
	//打印出九九乘法表
	//i 表示层数
	//两层循环和两个变量分别控制行和对应行输出的内容
	var n int =9
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print(j,"*",i,"=",i*j,"\t")
		}
		fmt.Println()
	}
}