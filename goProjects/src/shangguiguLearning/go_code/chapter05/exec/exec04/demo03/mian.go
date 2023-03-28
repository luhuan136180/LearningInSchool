package main
import "fmt"

func main()  {
	var month int
	fmt.Println("请输入月份:")
	fmt.Scanln(&month)
	switch month{
		case 3,4,5:
			fmt.Println("春季")
		case 6,7,8:
			fmt.Println("夏季")
		case 9,10,11:
			fmt.Println("秋季")
		case 12,1,2:
			fmt.Println("冬季")
		default:
			fmt.Println("输入有误")
	}
}
