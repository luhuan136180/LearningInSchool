package main

import "fmt"

func BubbleSort(arr *[7]int) {
	fmt.Println("arr=",*arr)
	temp:=0

	for i:=0;i<len(*arr);i++{//轮次
		num :=0
		for j:=0;j<len(*arr)-1-i;j++{//比较次数
			if (*arr)[j]>(*arr)[j+1]{
				//交换
				num++
				temp = (*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=temp
			}

			fmt.Printf("第%v轮的第%v次比较：%p \n",i+1,j+1,*arr)
		}
		if num==0{
			fmt.Println("不交换")
			break
		}
	}
	fmt.Println("arr=",*arr)
}

func main() {
	var arr [7]int=[7]int{5,6,3,24,5,3,4}
	BubbleSort(&arr)
	fmt.Println("main arr=", arr) //有序? 是有序的
}
