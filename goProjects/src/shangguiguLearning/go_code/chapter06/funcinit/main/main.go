package main

import (
	"fmt"
	"shangguiguLearning/go_code/chapter06/funcinit/utils"
)
var age = test()//定义全局变量
func test() int {
	fmt.Println("test()....")
	return 90
}
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()....age=",age)
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}