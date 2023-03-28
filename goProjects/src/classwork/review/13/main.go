package main

import "fmt"

func main() {
	var arr = [10]int{1, 23, 45, 12, 13, 12, 23, 42, 11, 17}
	fmt.Println(find(arr))
}
func find(arr [10]int) (max int, min int) {
	max, min = 0, 0
	for i, v := range arr {
		if arr[min] > v {
			min = v
		}
		if arr[max] < v {
			max = i
		}
	}
	return max, min
}
