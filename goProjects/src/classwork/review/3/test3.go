package main

import "fmt"

func main() {
	//for i := 100; i <= 1000; i++ {
	//	a:=i/100
	//	b:=i/10%10
	//	c:=i%10
	//}
	n := 123
	a := n / 100
	b := n / 10 % 10
	c := n % 10
	fmt.Println(a, b, c)
}
