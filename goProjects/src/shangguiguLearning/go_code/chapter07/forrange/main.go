package main

import "fmt"

func main() {
	heros := [...]string{"songjiang","aaaa","sda fafa"}
	for i , v :=range heros{
		fmt.Printf("i=%v  ,v=%v  ",i,v)
	}
	for _, v := range heros {
		fmt.Printf("元素的值=%v\n", v)
	}
}
