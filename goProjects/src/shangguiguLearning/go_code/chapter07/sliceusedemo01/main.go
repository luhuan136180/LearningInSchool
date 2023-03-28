package main

import "fmt"

func main() {
	var slice01 []float64 = make([]float64, 5, 10)
	slice01[1] = 100
	slice01[2] = 23
	fmt.Println("slice=", slice01)
	fmt.Println("slice 的size=",len(slice01))
	fmt.Println("slice 的cap=",cap(slice01))

	//方式3
	fmt.Println()
	//第3种方式：定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var strSlice []string = []string{"tom","jack","alice"}
	fmt.Println("strslice=",strSlice)
	fmt.Println("strslice size=",len(strSlice))
	fmt.Println("strslice cap=",cap(strSlice))
}
