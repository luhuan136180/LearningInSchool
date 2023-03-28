package main
import "fmt"

func main(){
	//使用while方式输出10句 "hello,world"
	var i int=1
	for{
		if i > 10{
			break
		}
		fmt.Println("hello,world",i)
		i++
	}
	fmt.Println("i=",i)

	//使用的do...while实现完成输出10句”hello,ok“
	var j int =1
	for{
		fmt.Println("HELLO,OK",j)
		j++
		if j>10 {
			break
		}
	}
}
