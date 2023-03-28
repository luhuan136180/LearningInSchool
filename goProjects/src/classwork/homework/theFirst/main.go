package main

import "fmt"

func main() {
	var n int
	fmt.Println("请输入数字n")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		j:=1+2*(i-1)
		k:=n-i
		for m := 1; m <= k; m++ {
			fmt.Print(" ")
		}
		for m:=1;m<=j;m++{
			fmt.Print("*")
		}
	fmt.Println()
	}
	for i := n-1; i >= 1; i-- {
		j:=1+2*(i-1)
		k:=n-i
		for m := 1; m <= k; m++ {
			fmt.Print(" ")
		}
		for m:=1;m<=j;m++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
