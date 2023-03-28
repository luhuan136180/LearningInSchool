package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Printf("please input a b c")
	fmt.Scanf("%d%d%d",&a,&b,&c)
	if a>c{
		a,c=c,a
	}
	if a>b{
		a,b=b,a
	}
	if b>c{
		b,c=c,b
	}
	fmt.Println("THE result is")
	fmt.Println("a=",a,"b=",b,"c=",c)
}
