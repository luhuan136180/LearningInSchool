package main

import (
	"fmt"
)

func main() {
	var num = []float64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}
	fmt.Println(average(num))
}
func average(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum / float64(len(slice))
}
