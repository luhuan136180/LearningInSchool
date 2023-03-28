package main

import "fmt"
//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}
func (p Phone) Calling() {
	fmt.Println("打电话....")
}

type Camera struct {
	name string
}
//让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}
type Computer struct {

}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	if Phone,ok :=usb.(Phone);ok{
		Phone.Calling()
	}
	usb.Stop()

}
func main() {
	var usbArr [3]Usb
	usbArr[0]=Phone{
		name: "vivo",
	}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	//遍历usbArr
	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb 接口声明的方法外，还需要调用Phone 特有方法 call. =》类型断言
	var computer Computer
	for _,v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}
	fmt.Println(usbArr)
}

