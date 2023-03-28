package main

import "fmt"

var(
	fun1 = func(n1 int,n2 int)int{
		return n1*n2
	}
)

func main() {
	res1 :=func(n1 int,n2 int)int{
		return n1+n2
	}(10,20)
	fmt.Println("res1=",res1)

	//将匿名函数func (n1 int, n2 int) int赋给 a变量
	//则a 的数据类型就是函数类型 ，此时,我们可以通过a完成调用
	a:=func (n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(90, 30)
	fmt.Println("res3=", res3)

	res4 := fun1(4,5)
	fmt.Println("res4=", res4)
}
