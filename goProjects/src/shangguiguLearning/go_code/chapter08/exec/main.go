package main

import (
	"fmt"
)

func main() {
	/*
		定义二维数组，用于保存三个班，每个班五名同学成绩，
		并求出每个班级平均分、以及所有班级平均分
	*/
	var arr [3][5]int
	for i:=0;i<len(arr);i++{
		fmt.Printf("开始录入%v班的成绩 \n",i+1)
		for j:=0;j<len(arr[i]);j++{
			fmt.Printf("请输入%v班的第%v学生的成绩 \n",i+1,j+1)
			fmt.Scanln(&arr[i][j])
		}

	}

	var allsum float64=0.0
	for i:=0;i<len(arr);i++{
		var sum float64=0.0
		for j :=0;j<len(arr[i]);j++{
			sum+=float64(arr[i][j])
		}
		allsum+=sum
		fmt.Printf("第%v班的平均成绩为：%v",i+1,sum/float64(len(arr[i])))
	}
	fmt.Printf("搜有班的平均成绩为：%v",allsum/float64(len(arr)*len(arr[0])))
}
