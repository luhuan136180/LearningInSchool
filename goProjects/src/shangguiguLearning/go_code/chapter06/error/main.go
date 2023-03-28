package main

import (
	"errors"
	"fmt"
)

//func test() {
//	//使用defer + recover 来捕获和处理异常
//	defer func(){
//		err := recover()// recover()内置函数，可以捕获到异常
//		if err != nil {// 说明捕获到错误
//			fmt.Println("err=",err)
//			fmt.Println("发送邮件给admin@sohu.com~")
//		}
//	}()
//	num1 :=10
//	num2 :=0
//	res := num1 / num2
//	fmt.Println("res= ",res)
//}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误")
	}

}

func test02() {
	err := readConf("config2.ini")
	if err != nil{
		//panic(err)
	}
	fmt.Println("test02()继续zhixing")
}
func main() {
	//test()
	//for {
	//	fmt.Println("main()下面的代码...")
	//	time.Sleep(time.Second)
	//}
	test02()
	fmt.Println("main()下面的代码...")
}
