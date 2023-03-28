package main

import "fmt"

func main() {
	var slice = []float64{1.0, 1.5, 2.2, 3.6}
	var sum float64
	for _, v := range slice {
		sum += v
	}
	avg := sum / float64(len(slice))
	fmt.Println("avg=", avg)
}
