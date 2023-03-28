package main

import (
	"fmt"
)
type Circle struct {
	radius float64
}

func (C Circle)Area()  float64{
	 return 3.14*C.radius*C.radius
}
//为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle)area2()float64  {
	//因为 c是指针，因此我们标准的访问其字段的方式是 (*c).radius
	//return 3.14 * (*c).radius * (*c).radius
	// (*c).radius 等价  c.radius
	fmt.Printf("c是*Cricle 指向的地址=%p",c)
	c.radius=10
	return 3.14*c.radius*c.radius
}
func main() {
	// 1)声明一个结构体Circle, 字段为 radius
	// 2)声明一个方法area和Circle绑定，可以返回面积。
	// 3)提示：画出area执行过程+说明
	var c1 Circle
	c1.radius=4.0
	res :=c1.Area()
	fmt.Println("res=",res)

	var c2 Circle
	fmt.Printf("main c 结构体变量地址 = %p \n",&c2)
	res2 := c2.area2()
	fmt.Println("面积=",res2)
	fmt.Println("c2.radius=",c2.radius)//10
}
