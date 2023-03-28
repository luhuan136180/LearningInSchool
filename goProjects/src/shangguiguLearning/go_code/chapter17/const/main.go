package main
import "fmt"
func main() {
	var num int
	num = 9 //ok
	//常量声明的时候，必须赋值。
	const tax int = 0
	//常量是不能修改
	//tax = 10
	fmt.Println(num, tax)
	//常量只能修饰bool、数值类型(int, float系列)、string 类型

	//fmt.Println(b)

	const (//表示a赋值0，其后的常量赋值依次加一
		a = iota//0//iota的本质是将行号转化为赋值
		b//1
		c//2
		d//3
	)


	fmt.Println(a, b, c, d)
}