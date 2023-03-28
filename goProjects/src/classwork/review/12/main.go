package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	fmt.Println("area=", area(a, b, c))
}
func area(a, c, b float64) float64 {
	var s = (a + b + c) / 2.0
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}
