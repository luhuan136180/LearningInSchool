package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point =Point{1,3}
	a=point//ok
	var b Point
	fmt.Println("B=",b)
	b = a.(Point)
	fmt.Println("b=",b)


	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x=b2
	y,ok :=x.(float32)
	if ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	}else {
		fmt.Println("convert fail")
	}
fmt.Println("继续执行")
}
