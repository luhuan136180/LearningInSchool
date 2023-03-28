package main

import "fmt"

func main() {
	var str string
	fmt.Println("请输入字符串")
	fmt.Scanln(&str)
	for _, v := range str {
		fmt.Println(string(v))
	}
}
