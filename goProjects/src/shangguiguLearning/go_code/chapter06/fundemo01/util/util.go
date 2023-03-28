package util

import "fmt"
//包中的函数名首字母大写（如：Cal）表示：函数Cal为公开函数=>>public
func Cal(n1 float64,n2 float64,operator byte) float64  {
	var res float64
	switch operator{
	case '+':
		res=n1+n2
	case '-':
		res=n1-n2
	case '*':
		res = n1*n2
	case '/':
		res=n1/n2
	default:
		fmt.Println("操作符号有误")
	}
	return res
}