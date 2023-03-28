package main

import "fmt"

func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int)  {
	if leftIndex>rightIndex{
		fmt.Println("没找到")
		return
	}
	var middle int = (leftIndex+rightIndex)/2
	if (*arr)[middle] ==findVal{
		fmt.Println("zhaodaol xiabiaowei= ,",middle)
	}else if (*arr)[middle]<findVal {
		BinaryFind(arr,leftIndex,middle-1,findVal)
	}else {
		BinaryFind(arr,middle+1,rightIndex,findVal)
	}
}
func main()  {
	arr :=[6]int{1,8,10,89,1000,23131}
	BinaryFind(&arr,0,len(arr)-1,10)
}
