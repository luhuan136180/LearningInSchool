package main

import "fmt"

func main() {
	var odd, even float64 = 0, 1
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even += float64(i)
		} else {
			odd += float64(i)
		}
	}
	fmt.Printf("odd=%v,even=%v", odd, even)
}
