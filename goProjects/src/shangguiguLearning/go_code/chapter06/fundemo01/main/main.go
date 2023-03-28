package main
import (
	"fmt"
	"shangguiguLearning/go_code/chapter06/fundemo01/util"
)


func main()  {
	var n1 float64 =1.2
	var n2 float64 =2.3
	var operator byte = '+'
	result := util.Cal(n1,n2,operator)
	fmt.Println("result=~~",result)
}

