package main
import "fmt"

func main() {
	num1:= 100
	fmt.Println("num1的类型%T,num1的值=%v,num1的地址为：",num1,num1,&num1)
	num2:=new(int)//*int
	*num2 = 100
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址%v\\n num2这个指针，指向的值=%v",
		num2, num2, &num2, *num2)
}