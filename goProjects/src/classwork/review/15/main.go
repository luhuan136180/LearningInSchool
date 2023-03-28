package main

import "fmt"

func main() {
	var slice []int
	slice = []int{1, 23, 4, 523, 4, 23, 2}
	fmt.Println(max(slice))
}
func max(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if max < v {
			max = v
		}
	}
	return max
}
