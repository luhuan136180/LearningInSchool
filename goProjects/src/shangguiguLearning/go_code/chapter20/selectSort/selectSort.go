package main

import (
	"fmt"
	"math/rand"
)

func selectSort(arr *[8]int) {
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		temp := min
		arr[minIndex] = arr[i]
		arr[i] = temp
		fmt.Println(arr)
	}
}
func main() {
	var arr [8]int
	for i := 0; i < 8; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	selectSort(&arr)
	fmt.Println(arr)
}
