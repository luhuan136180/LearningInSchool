package main

import (
	"fmt"
)

func main() {
	var arr [2][3]int=  [2][3]int{{1,2,3}, {4,10,6}}
	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr[0]);j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

////for-range来遍历二维数组
	for i,v := range arr{
		for j,v2 :=range v{
			fmt.Printf("arr[%v][%v]=%v \t",i,j,v2)
		}
		fmt.Println()
	}
}
