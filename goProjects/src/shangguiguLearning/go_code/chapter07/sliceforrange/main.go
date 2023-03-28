package main

import "fmt"

func main() {
	var arr [5]int=[...]int{1,2,3,4,5}
	slice := arr[1:4]
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v \n",i,slice[i])
	}
	//
	for index,value :=range slice{
		fmt.Printf("slice[%v]=%v",index,value)
	}
	fmt.Println("")
	//
	slice2 := slice[1:2]
	slice2[0]=3000// 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化
	fmt.Println("slice2= ",slice2)
	fmt.Println("slice=",slice)
	fmt.Println("arr=",arr)
	fmt.Println("")
	////用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100,200,300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3,400,500,600)
	fmt.Println("slice3",slice3)
	//通过append将切片slice3追加给slice3
	slice3= append(slice3,slice3...)
	fmt.Println("slice3",slice3)
	fmt.Println("")
	//	//切片的拷贝操作
	//	//切片使用copy内置函数完成拷贝，举例说明
	var slice4 []int = []int{1,2,3,4,5}
	var slice5 =make([]int,10)
	copy(slice5,slice4)
	fmt.Println("slice4=",slice4)
	fmt.Println("slice5=",slice5)

}
