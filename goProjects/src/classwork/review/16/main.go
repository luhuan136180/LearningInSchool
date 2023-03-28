package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fac(n))
}
func fac(n int) int {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}
