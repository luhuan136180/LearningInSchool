package main

import "fmt"
func fbn(n int)([]int64){
	//
	var slice =make([]int64,n)
	slice[0]=1
	slice[1]=1
	for i:=2;i<n;i++{
		slice[i]=slice[i-1]+slice[i-2]
	}
	return slice
}
func main()  {
	slice:=fbn(10)
	fmt.Println("slice:",slice)
}
