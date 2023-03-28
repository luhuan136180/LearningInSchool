package main

import "fmt"

func init() {
	fmt.Println("AAAA")
}
func main() {
	var arr = [10]int{231, 213, 212, 431, 123, 312, 221, 123, 111, 222}
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	fmt.Println("max=", max, "min=", min)
}
