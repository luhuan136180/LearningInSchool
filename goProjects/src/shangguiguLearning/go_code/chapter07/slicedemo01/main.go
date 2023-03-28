package main

import "fmt"

func main()  {
	var intArr [5]int =[5]int{1,22,33,66,99}
	//声明/定义一个切片
	//slice := intArr[1:3]
	//1. slice 就是切片名
	//2. intArr[1:3] 表示 slice 引用到intArr这个数组
	//3. 引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)

	//第一种方式：定义一个切片，然后让切片
	// 去引用一个已经创建好的数组，比如前面的案例就是这
	//样的。
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice) //  22, 33
	fmt.Println("slice 的元素个数 =", len(slice)) // 2
	fmt.Println("slice 的容量 =", cap(slice)) // 切片的容量是可以动态变化

}
