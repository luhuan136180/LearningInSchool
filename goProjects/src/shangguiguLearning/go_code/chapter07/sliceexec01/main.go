package main

import "fmt"

func test(slice []int) {
	slice[0]=200
}
func main() {
	var arr [5]int=[5]int{1,2,3,4,5}
	var slice =arr[:]
	var slice2=slice
	slice2[0]=100
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("arr=",arr)
	test(slice2)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
	fmt.Println("arr=",arr)

}
