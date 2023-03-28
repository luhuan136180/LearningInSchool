package main

import "fmt"
//定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
//struct 一种自定义的数据类型,值类型
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
	Scores [3]int
}
func main() {
	// 张老太养了20只猫猫:一只名字叫小白,今年3岁,白色。还有一只叫小花,
	// 今年100岁,花色。请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，
	// 年龄，颜色。如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。

	// //1. 使用变量的处理
	// var cat1Name string = "小白"
	// var cat1Age int = 3
	// var cat1Color string = "白色"

	// var cat2Name string = "小花"
	// var cat2Age int = 100
	// var cat2Color string = "花色"

	// //2. 使用数组解决
	// var catNames [2]string = [...]string{"小白", "小花"}
	// var catAges [2]int = [...]int{3, 100}
	// var catColors [2]string = [...]string{"白色", "花色"}
	// //... map[string]string

	// fmt.Println("ok")
	// 使用struct来完成案例
	var cat1 Cat
	fmt.Println(cat1)
	fmt.Printf("cat1的地址=%p\n", &cat1)
	cat1.Name="jsjs"
	cat1.Age=12
	cat1.Color="白色"
	cat1.Hobby="chi"
	fmt.Println("cat1=",cat1)
	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)

}
